package resolvers

import (
	"github.com/redis/go-redis/v9"
	"server/internal/pkg/misc"
	"server/pkg/ogm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	OGM   *ogm.OGM
	Cache *redis.Client
	Env   *misc.ENV
}
