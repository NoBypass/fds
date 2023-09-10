package generated

// Code automatically generated; DO NOT EDIT.

import "github.com/graphql-go/graphql"

var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"account": AccountQuery,
			"discord": DiscordQuery,
			"player":  PlayerQuery,
			"apiKey":  ApiKeyQuery,
		},
	},
)

var rootMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"signin":        SigninMutation,
			"createDiscord": CreateDiscordMutation,
			"giveXp":        GiveXpMutation,
		},
	},
)
var RootSchema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	},
)
