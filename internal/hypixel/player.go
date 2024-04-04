package hypixel

import (
	"encoding/json"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/opentracing/opentracing-go"
	"time"
)

func (c *APIClient) Player(name string, sp opentracing.Span) (*model.HypixelPlayerResponse, error) {
	url := "/player?name=" + name

	cached, ok := c.cache.Get(url)
	if ok {
		return cached.(*model.HypixelPlayerResponse), nil
	}

	body, err := c.Request(url, sp)
	if err != nil {
		return nil, err
	}

	var player model.HypixelPlayerResponse
	err = json.NewDecoder(body).Decode(&player)
	if err != nil {
		return nil, err
	}

	c.cache.Set(url, &player, 3*time.Minute)
	return &player, nil
}
