package schemas

import (
	"github.com/graphql-go/graphql"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/src/db/models"
	"server/src/db/repository"
)

var accountType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Account",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
			"joined_at": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var signinType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Signin",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
			"joined_at": &graphql.Field{
				Type: graphql.String,
			},
			"token": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var AccountQueryByUsername = &graphql.Field{
	Type: accountType,
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return repository.FindAccountByName(p.Context, p.Context.Value("driver").(neo4j.DriverWithContext), p.Args["username"].(string))
	},
}

var Signin = &graphql.Field{
	Type: signinType,
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
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
		remember, ok := p.Args["remember"].(bool)
		if !ok {
			remember = false
		}

		return repository.Signin(p.Context, p.Context.Value("driver").(neo4j.DriverWithContext), &models.AccountDto{
			Username: p.Args["username"].(string),
			Password: p.Args["password"].(string),
			Remember: remember,
		})
	},
}

// var ResetPassword = &graphql.Field{
// 	Type: accountType,
// 	Args: graphql.FieldConfigArgument{
// 		"password": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(graphql.String),
// 		},
// 	},
// 	Resolve: // TODO
// }
