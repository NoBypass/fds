package generated

// Code automatically generated; DO NOT EDIT.

import (
	"github.com/graphql-go/graphql"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/src/graph/services"
)

type Signin struct {
	Token   string  `json:"token"`
	Role    string  `json:"role"`
	Account Account `json:"account"`
}

var SigninType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Signin", Fields: graphql.Fields{
		"token": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"role": &graphql.Field{
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

	role, err := neo4j.GetProperty[string](r, "role")
	if err != nil {
		return nil, err
	}

	account, err := ResultToAccount(result)
	if err != nil {
		return nil, err
	}

	return &Signin{
		Token:   token,
		Role:    role,
		Account: *account,
	}, nil
}

type Account struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

var AccountType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Account", Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.Field{
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

	password, err := neo4j.GetProperty[string](r, "password")
	if err != nil {
		return nil, err
	}

	createdAt, err := neo4j.GetProperty[string](r, "created_at")
	if err != nil {
		return nil, err
	}

	return &Account{
		Name:      name,
		Password:  password,
		CreatedAt: createdAt,
	}, nil
}

type ApiKeyInput struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

var ApiKeyQuery = &graphql.Field{
	Type: SigninType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"role": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		input := &ApiKeyInput{
			Name: p.Args["name"].(string),
			Role: p.Args["role"].(string)}

		return services.ApiKeyQuery(p.Context, input), nil
	},
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

		return services.SigninMutation(p.Context, input), nil
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

		return services.AccountQuery(p.Context, input), nil
	},
}
