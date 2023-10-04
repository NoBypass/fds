package generated

// Code automatically generated; DO NOT EDIT.

import (
	"github.com/graphql-go/graphql"
	"server/internal/app/resolvers"
	"server/internal/pkg/generated/models"
)

var discordType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Discord", Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"lastDailyAt": &graphql.Field{
			Type: graphql.Int,
		},
		"discordId": &graphql.Field{
			Type: graphql.String,
		},
		"streak": &graphql.Field{
			Type: graphql.Int,
		},
		"level": &graphql.Field{
			Type: graphql.Int,
		},
		"xp": &graphql.Field{
			Type: graphql.Int,
		},
	},
},
)
var DiscordQuery = &graphql.Field{
	Type: discordType,
	Args: graphql.FieldConfigArgument{
		"discordId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		input := &models.DiscordInput{
			DiscordId: p.Args["discordId"].(string)}

		return resolvers.DiscordQuery(p.Context, input)
	},
}

var CreateDiscordMutation = &graphql.Field{
	Type: discordType,
	Args: graphql.FieldConfigArgument{
		"discordId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		input := &models.CreateDiscordInput{
			DiscordId: p.Args["discordId"].(string),
			Name:      p.Args["name"].(string)}

		return resolvers.CreateDiscordMutation(p.Context, input)
	},
}

var GiveXpMutation = &graphql.Field{
	Type: discordType,
	Args: graphql.FieldConfigArgument{
		"discordId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"amount": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		input := &models.GiveXpInput{
			DiscordId: p.Args["discordId"].(string),
			Amount:    p.Args["amount"].(int64)}

		return resolvers.GiveXpMutation(p.Context, input)
	},
}
