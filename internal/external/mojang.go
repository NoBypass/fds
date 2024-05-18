package external

import (
	"encoding/json"
	"fmt"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/NoBypass/fds/internal/pkg/utils"
	"github.com/NoBypass/mincache"
	"github.com/opentracing/opentracing-go"
	"io"
	"net/http"
	"time"
)

type MojangAPIClient struct {
	cache *mincache.Cache
}

func NewMojangAPIClient(cache *mincache.Cache) *MojangAPIClient {
	return &MojangAPIClient{
		cache: cache,
	}
}

func (c *MojangAPIClient) Request(name string, sp opentracing.Span) (io.ReadCloser, error) {
	var csp opentracing.Span
	if sp != nil {
		csp = opentracing.StartSpan("Mojang API", opentracing.ChildOf(sp.Context()))
		defer csp.Finish()

		csp.LogKV("name", name)
	}

	req, err := http.NewRequest(http.MethodGet, "https://api.mojang.com/users/profiles/minecraft/"+name, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("mojang: %s", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, &utils.ErrMojangAPINotFound{
			Message:  "mojang: player not found",
			Code:     http.StatusNotFound,
			Internal: err,
		}
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("mojang: unknown error")
	}

	return resp.Body, nil
}

func (c *MojangAPIClient) Player(name string, sp opentracing.Span) (*model.MojangProfile, error) {
	cached, ok := c.cache.Get("mojang:" + name)
	if ok {
		return cached.(*model.MojangProfile), nil
	}

	body, err := c.Request(name, sp)
	if err != nil {
		return nil, err
	}

	var player model.MojangProfile
	err = json.NewDecoder(body).Decode(&player)
	if err != nil {
		return nil, err
	}

	c.cache.Set(name, &player, 5*time.Minute)
	return &player, nil
}
