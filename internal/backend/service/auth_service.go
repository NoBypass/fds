package service

import (
	"fmt"
	"github.com/NoBypass/fds/internal/pkg/model"
	"github.com/NoBypass/fds/internal/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/opentracing/opentracing-go"
	"time"
)

type AuthService interface {
	Service
	SignJWT(<-chan *model.JWTClaims) <-chan string
	BotClaims(string) <-chan *model.JWTClaims
	BotPwdIsValid(string) bool
}

type authService struct {
	service
	cfg *utils.Config
}

func NewAuthService(cfg *utils.Config) AuthService {
	return &authService{
		cfg: cfg,
	}
}

func (s *authService) BotPwdIsValid(pwd string) bool {
	return pwd == s.cfg.BotPwd
}

func (s *authService) SignJWT(claimsCh <-chan *model.JWTClaims) <-chan string {
	out := make(chan string)

	s.Pipeline(func(func() opentracing.Span) error {
		claims := <-claimsCh

		var regClaims jwt.RegisteredClaims
		regClaims.Issuer = fmt.Sprintf("fds core %s", utils.VERSION)
		regClaims.IssuedAt = jwt.NewNumericDate(time.Now())
		if claims.Exp.After(time.Now()) {
			regClaims.ExpiresAt = jwt.NewNumericDate(claims.Exp)
		}
		regClaims.Subject = claims.Sub
		regClaims.Audience = func() []string {
			aud := make([]string, len(claims.Aud))
			for i, a := range claims.Aud {
				aud[i] = fmt.Sprintf("%d", a)
			}
			return aud
		}()

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, regClaims)
		signed, err := token.SignedString([]byte(s.cfg.JWTSecret))
		if err != nil {
			return err
		}

		out <- signed
		close(out)
		return nil
	}, s.SignJWT)

	return out
}

func (s *authService) BotClaims(sub string) <-chan *model.JWTClaims {
	out := make(chan *model.JWTClaims)

	s.Pipeline(func(func() opentracing.Span) error {
		claims := &model.JWTClaims{
			Sub: sub,
			Aud: []model.AuthRole{model.RoleBot},
		}

		out <- claims
		close(out)
		return nil
	}, s.BotClaims)

	return out
}
