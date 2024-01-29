package auth

type Service struct {
	secret string
}

func NewService(secret string) *Service {
	return &Service{secret: secret}
}
