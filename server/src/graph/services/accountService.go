package services

import (
	"context"
	"errors"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"net/http"
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

	claims, err := handlers.NewClaims("user")
	if err != nil {
		return nil, err
	}
	token, err := claims.Sign(account).Generate()

	return &generated.Signin{
		Token:   token,
		Role:    claims.Role,
		Account: *account,
	}, nil
}

func ApiKeyMutation(ctx context.Context, input *generated.ApiKeyInput) (*generated.Signin, error) {
	claims, err := handlers.ParseJWT(ctx.Value("request").(*http.Request).Header.Get("Authorization"))
	if err != nil {
		return nil, err
	}

	if claims.Role != "admin" {
		http.Error(ctx.Value("w").(http.ResponseWriter), "you need admin permissions to use chmod", http.StatusUnauthorized)
		return nil, errors.New("permission denied")
	}

	result, err := repository.FindAccountByName(ctx, ctx.Value("driver").(neo4j.DriverWithContext), input.Name)
	if err != nil {
		return nil, err
	}

	if result.Records == nil || len(result.Records) == 0 {
		return nil, errors.New("account not found")
	}

	account, err := generated.ResultToAccount(result)
	if err != nil {
		return nil, err
	}

	claims, err = handlers.NewClaims(input.Role)
	if err != nil {
		return nil, err
	}
	token, err := claims.Sign(account).Generate()

	return &generated.Signin{
		Token:   token,
		Role:    claims.Role,
		Account: *account,
	}, nil
}
