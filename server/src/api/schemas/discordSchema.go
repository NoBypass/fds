package schemas

import (
	"github.com/graphql-go/graphql"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/src/db/models"
	"server/src/db/repository"
)

type Discord struct {
	UUID        string  `json:"uuid"`
	DiscordID   string  `json:"discord_id"`
	Name        string  `json:"name"`
	Level       *int    `json:"level"`
	Xp          *int    `json:"xp"`
	Streak      *int    `json:"streak"`
	LastDailyAt *string `json:"last_daily_at"`
}

var discordType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Discord",
		Fields: graphql.Fields{
			"uuid": &graphql.Field{
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
			"last_daily_at": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DiscordQuery = &graphql.Field{
	Type: discordType,
	Args: graphql.FieldConfigArgument{
		"discord_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return repository.DiscordQuery(p.Context, p.Context.Value("driver").(neo4j.DriverWithContext), p.Args["discord_id"].(string))
	},
}

var CreateDiscordMutation = &graphql.Field{
	Type: discordType,
	Args: graphql.FieldConfigArgument{
		"discord_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return repository.CreateDiscordMutation(p.Context, p.Context.Value("driver").(neo4j.DriverWithContext), &models.DiscordDto{
			DiscordID: p.Args["discord_id"].(string),
			Name:      p.Args["name"].(string),
		})
	},
}
