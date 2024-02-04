package service

import (
	"github.com/NoBypass/fds/internal/app/repository"
	"golang.org/x/net/context"
)

type Service interface {
	InjectErrorChan() <-chan error
	InjectContext(ctx context.Context) context.CancelFunc
	Pipeline(run func(), repos ...repository.Repository)
}

type service struct {
	errCh chan<- error
	ctx   context.Context
}

func (s *service) InjectErrorChan() <-chan error {
	ch := make(chan error)
	s.errCh = ch
	return ch
}

func (s *service) InjectContext(ctx context.Context) context.CancelFunc {
	ctx, cancel := context.WithCancel(ctx)
	s.ctx = ctx
	return cancel
}

func (s *service) Pipeline(run func(), repos ...repository.Repository) {
	if repos != nil {
		for _, repo := range repos {
			repo.InjectContext(s.ctx)
		}
	}

	if s.ctx == nil {
		go run()
		return
	}

	select {
	case <-s.ctx.Done():
		return
	default:
		go run()
	}
}
