package service

type Service interface {
	InjectErrorChan() <-chan error
}

type service struct {
	errCh chan<- error
}
