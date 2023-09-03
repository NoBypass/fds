package generated

// Code automatically generated; DO NOT EDIT.

import (
	"github.com/graphql-go/graphql"
	"server/src/graph/generated/models"
	"server/src/graph/services"
)

var DiscordType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Discord", Fields: graphql.Fields{
		"discordId": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"level": &graphql.Field{
			Type: graphql.Int,
		},
		"xp": &graphql.Field{
			Type: graphql.Int,
		},
		"streak": &graphql.Field{
			Type: graphql.Int,
		},
		"lastDailyAt": &graphql.Field{
			Type: graphql.Int,
		},
		//"linkedWith": &graphql.Field{
		//	Type: LinkedWithType,
		//},
	},
},
)

//	var LinkedWithType = graphql.NewObject(graphql.ObjectConfig{
//		Name: "LinkedWith", Fields: graphql.Fields{
//			"linkedAt": &graphql.Field{
//				Type: graphql.NewNonNull(graphql.Int),
//			},
//			"player": &graphql.Field{
//				Type: graphql.NewNonNull(PlayerType),
//			},
//			"discord": &graphql.Field{
//				Type: graphql.NewNonNull(DiscordType),
//			},
//		},
//	},
//
// )
var DiscordQuery = &graphql.Field{
	Type: DiscordType,
	Args: graphql.FieldConfigArgument{
		"discordId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		input := &models.DiscordInput{
			DiscordId: p.Args["discordId"].(string)}

		return services.DiscordQuery(p.Context, input)
	},
}

var CreateDiscordMutation = &graphql.Field{
	Type: DiscordType,
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

		return services.CreateDiscordMutation(p.Context, input)
	},
}

var GiveXpMutation = &graphql.Field{
	Type: DiscordType,
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

		return services.GiveXpMutation(p.Context, input)
	},
}

//var CreateDiscordSubscription = &graphql.Field{
//	Type: DiscordType,
//	Args: graphql.FieldConfigArgument{
//		"discordId": &graphql.ArgumentConfig{
//			Type: graphql.NewNonNull(graphql.String),
//		},
//		"name": &graphql.ArgumentConfig{
//			Type: graphql.NewNonNull(graphql.String),
//		},
//	},
//	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
//		input := &models.CreateDiscordInput{
//			DiscordId: p.Args["discordId"].(string),
//			Name: p.Args["name"].(string),		}
//
//		return services.CreateDiscordSubscription(p.Context, input)
//	},
//}
//
//var GiveXpSubscription = &graphql.Field{
//	Type: DiscordType,
//	Args: graphql.FieldConfigArgument{
//		"discordId": &graphql.ArgumentConfig{
//			Type: graphql.NewNonNull(graphql.String),
//		},
//		"amount": &graphql.ArgumentConfig{
//			Type: graphql.NewNonNull(graphql.Int),
//		},
//	},
//	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
//		input := &models.GiveXpInput{
//			DiscordId: p.Args["discordId"].(string),
//			Amount: p.Args["amount"].(int64),		}
//
//		return services.GiveXpSubscription(p.Context, input)
//	},
//}
//
