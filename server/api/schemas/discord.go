package schemas

import (
	"github.com/graphql-go/graphql"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/db/models"
	"server/db/repository"
)

var DiscordType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Discord",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"discord_id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
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
			"last_daily_at": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DiscordQueryByDiscordId = &graphql.Field{
	Type: DiscordType,
	Args: graphql.FieldConfigArgument{
		"discord_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return repository.FindDiscordByDiscordId(p.Context, p.Context.Value("driver").(neo4j.DriverWithContext), p.Args["discord_id"].(string))
	},
}

var CreateDiscord = &graphql.Field{
	Type: DiscordType,
	Args: graphql.FieldConfigArgument{
		"discord_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return repository.CreateDiscord(p.Context, p.Context.Value("driver").(neo4j.DriverWithContext), &models.DiscordDto{
			DiscordID: p.Args["discord_id"].(string),
			Name:      p.Args["name"].(string),
		})
	},
}
