package service

type Service interface {
	InjectErrorChan() <-chan error
}

type service struct {
	errCh chan<- error
}

func (s *service) InjectErrorChan() <-chan error {
	ch := make(chan error)
	s.errCh = ch
	return ch
}
