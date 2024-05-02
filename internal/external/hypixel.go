package external

import (
	"encoding/json"
	"fmt"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/NoBypass/mincache"
	"github.com/labstack/gommon/log"
	"github.com/opentracing/opentracing-go"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type HypixelAPIClient struct {
	sync.Mutex
	cache     *mincache.Cache
	apiKey    string
	rateLimit int
	remaining int
	resetAt   time.Time
}

func NewHypixelAPIClient(cache *mincache.Cache, key string) *HypixelAPIClient {
	if key == "" {
		log.Fatal("hypixel: missing api key")
	}

	client := &HypixelAPIClient{
		cache:  cache,
		apiKey: key,
	}

	_, err := client.Request("/status?uuid=b876ec32e396476ba1158438d83c67d4", nil)
	if err != nil {
		log.Fatalf("unable to initialize hypixel client: %s", err)
	}

	return client
}

func (c *HypixelAPIClient) Request(url string, sp opentracing.Span) (io.ReadCloser, error) {
	var csp opentracing.Span
	if sp != nil {
		csp = opentracing.StartSpan("Hypixel API", opentracing.ChildOf(sp.Context()))
		defer csp.Finish()

		csp.LogKV(
			"url", url,
			"rateLimit", c.rateLimit,
			"remaining", c.remaining,
			"resetAt", c.resetAt.String(),
		)
	}

	if c.remaining > 0 && c.remaining < 10 {
		return nil, fmt.Errorf("hypixel: rate limited, reset in %s", c.resetAt)
	}

	req, err := http.NewRequest(http.MethodGet, "https://api.hypixel.net/v2"+url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("API-Key", c.apiKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("hypixel: %s", err)
	}

	c.parseRateLimit(resp.Header)

	if resp.StatusCode != http.StatusOK {
		var e model.HypixelError
		err = json.NewDecoder(resp.Body).Decode(&e)
		if err != nil {
			return nil, fmt.Errorf("hypixel: %s", resp.Status)
		}
		return nil, fmt.Errorf("hypixel: %s", e.Cause)
	}

	return resp.Body, nil
}

func (c *HypixelAPIClient) parseRateLimit(header http.Header) {
	rl, _ := strconv.Atoi(header.Get("RateLimit-Limit"))
	r, _ := strconv.Atoi(header.Get("RateLimit-Remaining"))
	reset, _ := strconv.Atoi(header.Get("RateLimit-Reset"))

	c.Lock()
	defer c.Unlock()

	c.rateLimit, c.remaining, c.resetAt = rl, r, time.Now().Add(time.Duration(reset)*time.Second)
}

func (c *HypixelAPIClient) Player(name string, sp opentracing.Span) (*model.HypixelPlayerResponse, error) {
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
