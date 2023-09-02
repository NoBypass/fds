package generated

// Code automatically generated; DO NOT EDIT.

import (
	"github.com/graphql-go/graphql"
	"server/src/graph/generated/models"
	"server/src/graph/services"
)

var PlayerType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Player", Fields: graphql.Fields{
		"UUID": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
},
)

var PlayerQuery = &graphql.Field{
	Type: PlayerType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		input := &models.PlayerInput{
			Name: p.Args["name"].(string)}

		return services.PlayerQuery(p.Context, input), nil
	},
}
