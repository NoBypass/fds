package adapter

import (
	"context"
	"fmt"
	"github.com/NoBypass/fds/internal/common"
	"github.com/NoBypass/fds/internal/domain"
	"github.com/NoBypass/mincache"
	"golang.org/x/time/rate"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type HypixelAPI struct {
	sync.Mutex

	cache *mincache.Cache
	api   *common.ExternalClient

	apiKey    string
	limiter   *rate.Limiter
	rateLimit int
	remaining int
	resetAt   time.Time
}

func NewHypixelAPI(cache *mincache.Cache, key string) *HypixelAPI {
	return &HypixelAPI{
		cache:  cache,
		apiKey: key,
		api:    common.NewExternalClient(cache, "https://api.hypixel.net", "Hypixel API"),
	}
}

func (c *HypixelAPI) Auctions(ctx context.Context, page int) (*domain.HypixelAuctionResponse, error) {
	auctions := new(domain.HypixelAuctionResponse)
	err := c.request(ctx, fmt.Sprintf("/skyblock/auctions?page=%d", page), false, &auctions)
	if err != nil {
		return nil, err
	}

	return auctions, nil
}

func (c *HypixelAPI) EndedAuctions(ctx context.Context) (*domain.HypixelAuctionResponse, error) {
	auctions := new(domain.HypixelAuctionResponse)
	err := c.request(ctx, "/skyblock/auctions_ended", false, &auctions)
	if err != nil {
		return nil, err
	}

	return auctions, nil
}

func (c *HypixelAPI) request(ctx context.Context, url string, rl bool, decode any) error {
	if c.remaining > 0 && c.remaining < 5 && rl {
		return fmt.Errorf("hypixel: rate limited, reset in %s", c.resetAt)
	}

	header, err := c.api.Request(ctx, url, time.Minute*5, decode)
	if err != nil {
		return err
	}

	if !rl {
		return nil
	}

	rlErr := c.parseRateLimit(header)
	if rlErr != nil || err != nil {
		return fmt.Errorf("hypixel: %w", err)
	}

	return nil
}

func (c *HypixelAPI) parseRateLimit(header *http.Header) error {
	rl, err := strconv.Atoi(header.Get("RateLimit-Limit"))
	if err != nil {
		return err
	}
	r, err := strconv.Atoi(header.Get("RateLimit-Remaining"))
	if err != nil {
		return err
	}
	reset, err := strconv.Atoi(header.Get("RateLimit-Reset"))
	if err != nil {
		return err
	}

	c.Lock()
	defer c.Unlock()

	c.rateLimit, c.remaining, c.resetAt = rl, r, time.Now().Add(time.Duration(reset)*time.Second)
	return nil
}
