package domain

import (
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

type JWTClaims struct {
	Aud []AuthRole
	Sub string
	Exp time.Time
}

func ParseAudience(aud []string) (AuthRole, error) {
	roles := make([]int, len(aud))
	for i, r := range aud {
		n, err := strconv.Atoi(r)
		if err != nil {
			return 0, err
		}

		roles[i] = n
	}

	smallest := roles[0]
	for _, r := range roles {
		if r < smallest {
			smallest = r
		}
	}
	return AuthRole(smallest), nil
}
