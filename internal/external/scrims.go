package external

import (
	"encoding/json"
	"fmt"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/NoBypass/mincache"
	"github.com/opentracing/opentracing-go"
	"io"
	"net/http"
	"time"
)

type ScrimsAPIClient struct {
	cache *mincache.Cache
}

func NewScrimsAPIClient(cache *mincache.Cache) *ScrimsAPIClient {
	return &ScrimsAPIClient{
		cache: cache,
	}
}

func (c *ScrimsAPIClient) Request(url string, sp opentracing.Span) (io.ReadCloser, error) {
	var csp opentracing.Span
	if sp != nil {
		csp = opentracing.StartSpan("Scrims API", opentracing.ChildOf(sp.Context()))
		defer csp.Finish()

		csp.LogKV("url", url)
	}

	req, err := http.NewRequest(http.MethodGet, "https://api.scrims.network/v1"+url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("scrims: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("scrims: unknown error")
	}

	return resp.Body, nil
}

func (c *ScrimsAPIClient) Player(name string, sp opentracing.Span) (*model.ScrimsPlayerResponse, error) {
	url := "/user?username=" + name

	cached, ok := c.cache.Get("scrims:" + url)
	if ok {
		return cached.(*model.ScrimsPlayerResponse), nil
	}

	body, err := c.Request(url, sp)
	if err != nil {
		return nil, err
	}

	var player model.ScrimsPlayerResponse
	err = json.NewDecoder(body).Decode(&player)
	if err != nil {
		return nil, err
	}

	c.cache.Set(url, &player, 5*time.Minute)
	return &player, nil
}
