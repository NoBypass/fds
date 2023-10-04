package middleware

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"net/http"
	"server/src/api/handlers"
	"server/src/auth"
	"time"
)

const SLIDING_WINDOW_DURATION = 60

func RateLimiterMiddleware(ctx context.Context) {
	claims, ok := ctx.Value("claims").(*auth.CustomClaims)
	var role string
	if ok {
		role = claims.Role
	}
	cache := ctx.Value("cache").(*redis.Client)
	req := ctx.Value("req").(*http.Request)
	res := ctx.Value("res").(*handlers.Responder)
	ip := req.RemoteAddr

	var rateLimit int64 = 25
	switch role {
	case ADMIN:
		rateLimit = 5000
	case BOT:
		rateLimit = 10000
	case SUBSCRIBER:
		rateLimit = 1000
	case USER:
		rateLimit = 250
	}

	now := time.Now().Unix()
	start := now - SLIDING_WINDOW_DURATION

	_, err := cache.ZAdd(ctx, ip, redis.Z{
		Score:  float64(now),
		Member: now,
	}).Result()
	if err != nil {
		res.Status(http.StatusInternalServerError)
		_ = res.AddError(err, handlers.KNOWN_ERROR, []string{"rateLimiter.go", "While adding timestamp to Redis"})
		return
	}

	cache.ZRemRangeByScore(ctx, ip, "-inf", fmt.Sprintf("(%d", start))

	numRequests, err := cache.ZCard(ctx, ip).Result()
	if err != nil {
		res.Status(http.StatusInternalServerError)
		_ = res.AddError(err, handlers.KNOWN_ERROR, []string{"rateLimiter.go", "While counting timestamps in Redis"})
		return
	}

	if numRequests > rateLimit {
		res.Status(http.StatusTooManyRequests)
		_ = res.AddError(fmt.Errorf("rate limit exceeded"), handlers.RATE_LIMIT_EXCEEDED, []string{"rateLimiter.go", "Rate limit exceeded"})
		return
	}
}
