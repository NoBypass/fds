package middleware

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"net/http"
	"server/internal/app/global"
	"server/internal/pkg/auth"
	"time"
)

const SLIDING_WINDOW_DURATION = 60

func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		claims, ok := ctx.Value("claims").(*auth.CustomClaims)
		var role string
		if ok {
			role = claims.Role
		}
		cache := global.Get().Cache
		ip := r.RemoteAddr

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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		cache.ZRemRangeByScore(ctx, ip, "-inf", fmt.Sprintf("(%d", start))

		numRequests, err := cache.ZCard(ctx, ip).Result()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if numRequests > rateLimit {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
