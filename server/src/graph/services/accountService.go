package services

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"net/http"
	"server/src/api/handlers"
	"server/src/graph/generated/models"
	"server/src/repository"
	"server/src/utils"
	"strconv"
	"time"
)

func AccountQuery(ctx context.Context, input *models.AccountInput) (*models.Account, error) {
	result, err := repository.FindAccountByName(ctx, ctx.Value("driver").(neo4j.DriverWithContext), input.Name)
	if err != nil {
		return nil, err
	}

	return models.ResultToAccount(result)
}

func SigninMutation(ctx context.Context, input *models.SigninInput) (*models.Signin, error) {
	result, err := repository.FindAccountByName(ctx, ctx.Value("driver").(neo4j.DriverWithContext), input.Name)
	if err != nil {
		return nil, err
	}

	if result.Records == nil || len(result.Records) == 0 {
		password, err := utils.Hash(input.Password)
		if err != nil {
			return nil, err
		}

		result, err = repository.CreateAccount(ctx, ctx.Value("driver").(neo4j.DriverWithContext), &models.Account{
			Name:      input.Name,
			Password:  password,
			CreatedAt: strconv.FormatInt(time.Now().Unix(), 10),
		})
		if err != nil {
			return nil, err
		}
	}

	account, err := models.ResultToAccount(result)
	if err != nil {
		return nil, err
	}

	ok := utils.CompareHash(input.Password, account.Password)
	if !ok {
		return nil, handlers.NewHttpError(ctx, http.StatusUnauthorized, "incorrect password")
	}

	claims, err := handlers.NewClaims("user")
	if err != nil {
		return nil, err
	}
	token, err := claims.Sign(account).Generate()

	return &models.Signin{
		Token:   token,
		Role:    claims.Role,
		Account: *account,
	}, nil
}

func ApiKeyQuery(ctx context.Context, input *models.ApiKeyInput) (*models.Signin, error) {
	claims, err := handlers.ParseJWT(ctx.Value("request").(*http.Request).Header.Get("Authorization"))
	if err != nil {
		return nil, err
	}

	if claims.Role != "admin" {
		return nil, handlers.NewHttpError(ctx, http.StatusUnauthorized, "you do not have permission to request a special api key")
	}

	result, err := repository.FindAccountByName(ctx, ctx.Value("driver").(neo4j.DriverWithContext), input.Name)
	if err != nil {
		return nil, err
	}

	handlers.CheckIfFound(ctx, result, "couldn't find account with name "+input.Name)

	account, err := models.ResultToAccount(result)
	if err != nil {
		return nil, err
	}

	claims, err = handlers.NewClaims(input.Role)
	if err != nil {
		return nil, err
	}
	token, err := claims.Sign(account).Generate()

	return &models.Signin{
		Token:   token,
		Role:    claims.Role,
		Account: *account,
	}, nil
}
