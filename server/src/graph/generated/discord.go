package generated

// Code automatically generated; DO NOT EDIT.

import (
	"github.com/graphql-go/graphql"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/src/graph/services"
)

type Discord struct {
	DiscordId   string `json:"discord_id"`
	Name        string `json:"name"`
	Level       int64  `json:"level"`
	Xp          int64  `json:"xp"`
	Streak      int64  `json:"streak"`
	LastDailyAt string `json:"last_daily_at"`
}

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
			Type: graphql.String,
		},
	},
},
)

func ResultToDiscord(result *neo4j.EagerResult) (*Discord, error) {
	r, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "d")
	if err != nil {
		return nil, err
	}

	discordId, err := neo4j.GetProperty[string](r, "discord_id")
	if err != nil {
		return nil, err
	}

	name, err := neo4j.GetProperty[string](r, "name")
	if err != nil {
		return nil, err
	}

	level, err := neo4j.GetProperty[int64](r, "level")
	if err != nil {
		return nil, err
	}

	xp, err := neo4j.GetProperty[int64](r, "xp")
	if err != nil {
		return nil, err
	}

	streak, err := neo4j.GetProperty[int64](r, "streak")
	if err != nil {
		return nil, err
	}

	lastDailyAt, err := neo4j.GetProperty[string](r, "last_daily_at")
	if err != nil {
		return nil, err
	}

	return &Discord{
		DiscordId:   discordId,
		Name:        name,
		Level:       level,
		Xp:          xp,
		Streak:      streak,
		LastDailyAt: lastDailyAt,
	}, nil
}

type DiscordInput struct {
	DiscordId string `json:"discordId"`
}

var DiscordQuery = &graphql.Field{
	Type: DiscordType,
	Args: graphql.FieldConfigArgument{
		"discordId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		input := &DiscordInput{
			DiscordId: p.Args["discordId"].(string)}

		return services.DiscordQuery(p.Context, input), nil
	},
}

type CreateDiscordInput struct {
	DiscordId string `json:"discordId"`
	Name      string `json:"name"`
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
		input := &CreateDiscordInput{
			DiscordId: p.Args["discordId"].(string),
			Name:      p.Args["name"].(string)}

		return services.CreateDiscordMutation(p.Context, input), nil
	},
}
