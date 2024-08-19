package domain

import (
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

type AuthRole int

const (
	RoleAdmin AuthRole = iota
	RoleBot
	RolePremium
	RoleMember
)

type JWTSigner struct {
	Secret        []byte
	DefaultClaims jwt.RegisteredClaims
}

func NewJWTSigner(secret string) *JWTSigner {
	return &JWTSigner{
		Secret: []byte(secret),
		DefaultClaims: jwt.RegisteredClaims{
			Issuer: "fds",
		},
	}
}

func (s JWTSigner) NewJWT(to string, role AuthRole) string {
	s.DefaultClaims.ExpiresAt = expireAt(role)
	s.DefaultClaims.Audience = role.toAud()
	s.DefaultClaims.Subject = to

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, s.DefaultClaims)
	signed, _ := token.SignedString(s.Secret)
	return signed
}

func expireAt(role AuthRole) *jwt.NumericDate {
	switch role {
	case RoleBot:
		return nil
	default:
		return jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour))
	}
}

func (r *AuthRole) toAud() []string {
	return []string{strconv.Itoa(int(*r))}
}

func CanAccess(aud []string, to AuthRole) bool {
	n, err := strconv.Atoi(aud[0])
	if err != nil {
		return false
	}
	return n <= int(to)
}
