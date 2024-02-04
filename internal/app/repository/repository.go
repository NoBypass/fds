package repository

import (
	"github.com/NoBypass/fds/internal/pkg/surreal_wrap"
	"golang.org/x/net/context"
)

type Repository interface {
	InjectContext(ctx context.Context)
}

type repository struct {
	*surreal_wrap.DB
}

func newRepository(db *surreal_wrap.DB) repository {
	return repository{
		DB: db,
	}
}
