package schemas

import (
	"github.com/graphql-go/graphql"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/db/models"
	"server/db/repository"
)

var AccountType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Account",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
			"joined_at": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var AccountQueryByUsername = &graphql.Field{
	Type: AccountType,
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return repository.FindAccountByName(p.Context, p.Context.Value("driver").(neo4j.DriverWithContext), p.Args["username"].(string))
	},
}

var CreateAccount = &graphql.Field{
	Type: AccountType,
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return repository.CreateAccount(p.Context, p.Context.Value("driver").(neo4j.DriverWithContext), &models.AccountDto{
			Username: p.Args["username"].(string),
			Password: p.Args["password"].(string),
		})
	},
}
