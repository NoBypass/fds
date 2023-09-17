package services

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"net/http"
	"server/src/api/handlers"
	"server/src/graph/generated/models"
	"server/src/repository"
	"server/src/repository/db"
	"server/src/utils"
	"strconv"
	"time"
)

func AccountQuery(ctx context.Context, input *models.AccountInput) (*models.Account, error) {
	result, err := repository.FindAccountByName(ctx, ctx.Value("driver").(neo4j.DriverWithContext), input.Name)
	if err != nil {
		return nil, err
	}

	return utils.MapResult(&models.Account{}, result, "a")
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

	account, err := utils.MapResult(&models.Account{}, result, "a")
	if err != nil {
		return nil, err
	}

	ok := utils.CompareHash(input.Password, account.Password)
	if !ok {
		return nil, handlers.NewHttpError(ctx, http.StatusUnauthorized, "incorrect password")
	}

	claims, err := utils.NewClaims("user")
	if err != nil {
		return nil, err
	}
	token, err := claims.Sign(account).Generate(ctx)

	return &models.Signin{
		Token:   token,
		Role:    claims.Role,
		Account: account,
	}, nil
}

func ApiKeyQuery(ctx context.Context, input *models.ApiKeyInput) (*models.Signin, error) {
	claims := ctx.Value("claims").(*utils.CustomClaims)
	err := utils.Allow(ctx, []string{"admin"})
	if err != nil {
		return nil, err
	}

	accounts := db.New[models.Account](ctx)
	account, err := accounts.Find(&models.Account{Name: input.Name})

	claims, err = utils.NewClaims(input.Role)
	if err != nil {
		return nil, err
	}
	token, err := claims.Sign(account).Generate(ctx)

	return &models.Signin{
		Token:   token,
		Role:    claims.Role,
		Account: account,
	}, nil
}
