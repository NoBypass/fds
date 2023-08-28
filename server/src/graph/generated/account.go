package generated

// Code automatically generated; DO NOT EDIT.

import (
	"github.com/graphql-go/graphql"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/src/repository"
)

type Signin struct {
	Token     string  `json:"token"`
	ExpiresAt string  `json:"expires_at"`
	Account   Account `json:"account"`
}

var SigninType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Signin", Fields: graphql.Fields{
		"token": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"expiresAt": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"account": &graphql.Field{
			Type: graphql.NewNonNull(AccountType),
		},
	},
},
)

func ResultToSignin(result *neo4j.EagerResult) (*Signin, error) {
	r, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "s")
	if err != nil {
		return nil, err
	}

	token, err := neo4j.GetProperty[string](r, "token")
	if err != nil {
		return nil, err
	}

	expiresAt, err := neo4j.GetProperty[string](r, "expires_at")
	if err != nil {
		return nil, err
	}

	account, err := ResultToAccount(result)
	if err != nil {
		return nil, err
	}

	return &Signin{
		Token:     token,
		ExpiresAt: expiresAt,
		Account:   *account,
	}, nil
}

type Account struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

var AccountType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Account", Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"role": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"createdAt": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
},
)

func ResultToAccount(result *neo4j.EagerResult) (*Account, error) {
	r, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "a")
	if err != nil {
		return nil, err
	}

	name, err := neo4j.GetProperty[string](r, "name")
	if err != nil {
		return nil, err
	}

	email, err := neo4j.GetProperty[string](r, "email")
	if err != nil {
		return nil, err
	}

	role, err := neo4j.GetProperty[string](r, "role")
	if err != nil {
		return nil, err
	}

	createdAt, err := neo4j.GetProperty[string](r, "created_at")
	if err != nil {
		return nil, err
	}

	return &Account{
		Name:      name,
		Email:     email,
		Role:      role,
		CreatedAt: createdAt,
	}, nil
}

type SigninInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Remember bool   `json:"remember"`
}

var SigninMutation = &graphql.Field{
	Type: SigninType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"remember": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		input := &SigninInput{
			Name:     p.Args["name"].(string),
			Password: p.Args["password"].(string),
			Remember: p.Args["remember"].(bool)}

		return repository.SigninMutation(&p.Context, input), nil
	},
}

type AccountInput struct {
	Name string `json:"name"`
}

var AccountQuery = &graphql.Field{
	Type: AccountType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		input := &AccountInput{
			Name: p.Args["name"].(string)}

		return repository.AccountQuery(&p.Context, input), nil
	},
}
