package generated

// Code automatically generated; DO NOT EDIT.

import (
	"github.com/graphql-go/graphql"
	"server/src/db/repository"
)

type Player struct {
	Uuid *string `json:"uuid"`
	Name *string `json:"name"`
}


var playerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Player",
		Fields: graphql.Fields{
			"uuid": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
)


var PlayerQuery = &graphql.Field{
	Type: playerType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return repository.PlayerQuery(p)
	},
}


