package generated

// Code automatically generated; DO NOT EDIT.

import (
	"github.com/graphql-go/graphql"
	"server/src/graph/generated/models"
	"server/src/graph/services"
)

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
		input := &models.ApiKeyInput{
			Name: p.Args["name"].(string),
			Role: p.Args["role"].(string)}

		return services.ApiKeyQuery(p.Context, input)
	},
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
		input := &models.SigninInput{
			Name:     p.Args["name"].(string),
			Password: p.Args["password"].(string),
			Remember: p.Args["remember"].(bool)}

		return services.SigninMutation(p.Context, input)
	},
}

var AccountQuery = &graphql.Field{
	Type: AccountType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		input := &models.AccountInput{
			Name: p.Args["name"].(string)}

		return services.AccountQuery(p.Context, input)
	},
}
