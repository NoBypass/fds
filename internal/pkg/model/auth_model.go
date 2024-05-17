package model

import "time"

type AuthRole int

const (
	RoleAdmin AuthRole = iota
	RoleBot
	RolePremium
	RoleMember
)

type JWTClaims struct {
	Aud []AuthRole
	Sub string
	Exp time.Time
}
