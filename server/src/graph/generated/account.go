package generated

// Code automatically generated; DO NOT EDIT.

import (
	"github.com/graphql-go/graphql"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/src/db/repository"
)

type Signin struct {
	Token     string  `json:"token"`
	ExpiresAt string  `json:"expires_at"`
	Account   account `json:"account"`
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
			Type: graphql.NewNonNull(graphql.Account),
		},
	},
},
)

func ResultToSignin(r *neo4j.EagerResult) (*Signin, error) {
	result, _, err := neo4j.GetRecordValue[neo4j.Node](r.Records[0], "%!s(uint8=115)")
	if err != nil {
		return nil, err
	}

	token, err := neo4j.GetProperty[string](result, "token")
	if err != nil {
		return nil, err
	}

	expiresAt, err := neo4j.GetProperty[string](result, "expires_at")
	if err != nil {
		return nil, err
	}

	account, err := neo4j.GetProperty[account](result, "account")
	if err != nil {
		return nil, err
	}

	return &Signin{
		Token:     token,
		ExpiresAt: expiresAt,
		Account:   account,
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

func ResultToAccount(r *neo4j.EagerResult) (*Account, error) {
	result, _, err := neo4j.GetRecordValue[neo4j.Node](r.Records[0], "%!s(uint8=97)")
	if err != nil {
		return nil, err
	}

	name, err := neo4j.GetProperty[string](result, "name")
	if err != nil {
		return nil, err
	}

	email, err := neo4j.GetProperty[string](result, "email")
	if err != nil {
		return nil, err
	}

	role, err := neo4j.GetProperty[string](result, "role")
	if err != nil {
		return nil, err
	}

	createdAt, err := neo4j.GetProperty[string](result, "created_at")
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

var AccountMutation = &graphql.Field{
	Type: SigninType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return repository.AccountMutation(p), nil
	},
}

var DiscordMutation = &graphql.Field{
	Type: SigninType,
	Args: graphql.FieldConfigArgument{
		"discordId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return repository.DiscordMutation(p), nil
	},
}

var PlayerMutation = &graphql.Field{
	Type: SigninType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return repository.PlayerMutation(p), nil
	},
}

var SigninMutation = &graphql.Field{
	Type: AccountType,
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
		return repository.SigninMutation(p), nil
	},
}

var CreateDiscordMutation = &graphql.Field{
	Type: AccountType,
	Args: graphql.FieldConfigArgument{
		"discordId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return repository.CreateDiscordMutation(p), nil
	},
}
