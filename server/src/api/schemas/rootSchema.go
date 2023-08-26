package schemas

import "github.com/graphql-go/graphql"

var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"account": AccountQueryByUsername,
			"discord": DiscordQueryByDiscordId,
			"player":  PlayerQueryByName,
		},
	},
)

var rootMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createDiscord": CreateDiscord,
			"signin":        Signin,
		},
	},
)

var RootSchema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	},
)
