package resolvers

import (
	"context"
	"server/internal/pkg/auth"
	"server/internal/pkg/generated/models"
	"server/internal/pkg/misc"
	"server/pkg/ogm"
	"strconv"
	"time"
)

func AccountQuery(ctx context.Context, input *models.AccountInput) (*models.Account, error) {
	accounts := ogm.New[models.Account](ctx)
	account, err := accounts.Find(&models.Account{Name: input.Name})
	if err != nil {
		return nil, err
	}

	if !auth.HasRole(ctx, "admin") {
		account.Password = ""
		account.CreatedAt = ""
	}

	return account, nil
}

func SigninMutation(ctx context.Context, input *models.SigninInput) (*models.Signin, error) {
	accounts := ogm.New[models.Account](ctx)
	account, err := accounts.Find(&models.Account{Name: input.Name})

	if err != nil {
		password, err := misc.Hash(input.Password)
		if err != nil {
			return nil, err
		}

		account, err = accounts.Create(&models.Account{
			Name:      input.Name,
			Password:  password,
			CreatedAt: strconv.FormatInt(time.Now().Unix(), 10),
		})
		if err != nil {
			return nil, err
		}
	}

	ok := misc.CompareHash(input.Password, account.Password)
	if !ok {
		return nil, err
	}

	claims, err := auth.NewClaims("user")
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
	err := auth.Allow(ctx, []string{"admin"})
	if err != nil {
		return nil, err
	}

	accounts := ogm.New[models.Account](ctx)
	account, err := accounts.Find(&models.Account{Name: input.Name})

	claims, err := auth.NewClaims(input.Role)
	if err != nil {
		return nil, err
	}
	token, err := claims.Sign(account).Generate(ctx)
	if err != nil {
		return nil, err
	}

	return &models.Signin{
		Token:   token,
		Role:    claims.Role,
		Account: account,
	}, nil
}
