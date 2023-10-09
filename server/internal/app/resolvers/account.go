package resolvers

import (
	"context"
	"server/internal/app/global"
	"server/internal/pkg/auth"
	"server/internal/pkg/generated/models"
	"server/internal/pkg/misc"
	"server/pkg/ogm"
	"strconv"
	"time"
)

func AccountQuery(ctx context.Context, input *models.AccountInput) (*models.Account, error) {
	records, err := global.Get().DB.Query("MATCH (a:Account { name: $name }) RETURN a", map[string]any{
		"name": input.Name,
	})
	if err != nil {
		return nil, err
	}
	account, err := ogm.Map(&models.Account{}, records, "a")
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
	db := global.Get().DB
	records, err := db.Query("MATCH (a:Account { name: $name }) RETURN a",
		map[string]any{
			"name": input.Name,
		})

	if err != nil {
		password, err := misc.Hash(input.Password)
		if err != nil {
			return nil, err
		}

		records, err = db.Query("CREATE (a:Account { name: $name, email: $email, password: $password, role: $role, created_at: $created_at }) RETURN a",
			map[string]any{
				"name":       input.Name,
				"password":   password,
				"created_at": strconv.FormatInt(time.Now().Unix(), 10),
			})
		if err != nil {
			return nil, err
		}
	}
	account, err := ogm.Map(&models.Account{}, records, "a")
	if err != nil {
		return nil, err
	}

	ok := misc.CompareHash(input.Password, account.Password)
	if !ok {
		return nil, err
	}

	claims, err := auth.NewClaims("user")
	if err != nil {
		return nil, err
	}
	token, err := claims.Sign(account).Generate()

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

	records, err := global.Get().DB.Query("MATCH (a:Account { name: $name }) RETURN a",
		map[string]any{
			"name": input.Name,
		})
	if err != nil {
		return nil, err
	}

	account, err := ogm.Map(&models.Account{}, records, "a")
	if err != nil {
		return nil, err
	}

	claims, err := auth.NewClaims(input.Role)
	if err != nil {
		return nil, err
	}
	token, err := claims.Sign(account).Generate()
	if err != nil {
		return nil, err
	}

	return &models.Signin{
		Token:   token,
		Role:    claims.Role,
		Account: account,
	}, nil
}
