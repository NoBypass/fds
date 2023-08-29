package services

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/src/api/handlers"
	"server/src/graph/generated"
	"server/src/repository"
	"server/src/utils"
	"strconv"
	"time"
)

func AccountQuery(ctx context.Context, input *generated.AccountInput) (*generated.Account, error) {
	result, err := repository.FindAccountByName(ctx, ctx.Value("driver").(neo4j.DriverWithContext), input.Name)
	if err != nil {
		return nil, err
	}

	return generated.ResultToAccount(result)
}

func SigninMutation(ctx context.Context, input *generated.SigninInput) (*generated.Signin, error) {
	result, err := repository.FindAccountByName(ctx, ctx.Value("driver").(neo4j.DriverWithContext), input.Name)
	if err != nil {
		return nil, err
	}

	if result.Records == nil || len(result.Records) == 0 {
		password, err := utils.Hash(input.Password)
		if err != nil {
			return nil, err
		}

		result, err = repository.CreateAccount(ctx, ctx.Value("driver").(neo4j.DriverWithContext), &generated.Account{
			Name:      input.Name,
			Password:  password,
			Role:      "user",
			CreatedAt: strconv.FormatInt(time.Now().Unix(), 10),
		})
		if err != nil {
			return nil, err
		}
	}

	account, err := generated.ResultToAccount(result)
	if err != nil {
		return nil, err
	}

	ok := utils.CompareHash(input.Password, account.Password)
	if !ok {
		return nil, errors.New("invalid password")
	}

	expiresAt := utils.GetNowInMs() + 60*60*24*7
	if input.Remember {
		expiresAt = utils.GetNowInMs() + 60*60*24*30
	}

	claims := handlers.CustomClaims{
		Username: account.Name,
		Role:     "user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  utils.GetNowInMs(),
		},
	}
	token, err := claims.Sign(account).Generate()

	return &generated.Signin{
		Token:   token,
		Account: *account,
	}, nil
}
