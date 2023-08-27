package repository

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/src/graph/generated"
	"server/src/utils"
	"strconv"
)

func FindAccountByName(ctx context.Context, driver neo4j.DriverWithContext, name string) (*generated.Account, error) { // TODO remove in order to use only FindAccountByToken
	result, err := neo4j.ExecuteQuery(ctx, driver,
		"MATCH (a:Account { username: $username }) RETURN a",
		map[string]any{
			"username": name,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	return generated.ResultToAccount(result)
}

func Signin(ctx context.Context, driver neo4j.DriverWithContext, accountDto *models.AccountDto) (*generated.Signin, error) {
	result, err := neo4j.ExecuteQuery(ctx, driver,
		"MATCH (a:Account { username: $username }) RETURN a",
		map[string]any{
			"username": accountDto.Username,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	if result.Records == nil || len(result.Records) == 0 {
		password, err := utils.Hash(accountDto.Password)
		if err != nil {
			return nil, err
		}

		result, err = neo4j.ExecuteQuery(ctx, driver,
			"CREATE (a:Account { username: $username, password: $password, joined_at: $joined_at }) RETURN a",
			map[string]any{
				"username":  accountDto.Username,
				"password":  password,
				"joined_at": strconv.FormatInt(utils.GetNowInMs(), 10),
			}, neo4j.EagerResultTransformer)
		if err != nil {
			return nil, err
		}
	}

	account, err := generated.ResultToAccount(result)
	if err != nil {
		return nil, err
	}

	ok := utils.CompareHash(accountDto.Password, account.Password)
	if !ok {
		return nil, errors.New("invalid password")
	}

	expiresAt := utils.GetNowInMs() + 60*60*24*7
	if accountDto.Remember {
		expiresAt = utils.GetNowInMs() + 60*60*24*30
	}

	token, err := utils.GenerateJWT(utils.CustomClaims{
		Username: account.Username,
		Role:     "user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  utils.GetNowInMs(),
		},
	})

	return &generated.Signin{
		ID:       account.ID,
		Username: account.Username,
		Password: account.Password,
		JoinedAt: account.JoinedAt,
		Token:    token,
	}, nil
}
