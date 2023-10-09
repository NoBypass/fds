package generated

// Code automatically generated; DO NOT EDIT.

import (
	"github.com/graphql-go/graphql"
	"server/internal/app/resolvers"
	"server/internal/pkg/generated/models"
)

var verifiedWithType = graphql.NewObject(graphql.ObjectConfig{
	Name: "VerifiedWith", Fields: graphql.Fields{
		"verifiedAt": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
},
)

var playerType = graphql.NewObject(graphql.ObjectConfig{
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
	Type: playerType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		input := &models.PlayerInput{
			Name: p.Args["name"].(string)}

		return resolvers.PlayerQuery(p.Context, input)
	},
}
