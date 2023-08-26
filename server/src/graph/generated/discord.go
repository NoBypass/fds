package generated

// Code automatically generated; DO NOT EDIT.

import (
	"github.com/graphql-go/graphql"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/src/db/repository"
)

type Discord struct {
	Uuid        string `json:"uuid"`
	Name        string `json:"name"`
	Level       int64  `json:"level"`
	Xp          int64  `json:"xp"`
	Streak      int64  `json:"streak"`
	LastDailyAt string `json:"last_daily_at"`
}

var DiscordType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Discord", Fields: graphql.Fields{
		"UUID": &graphql.Field{
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

	UUID, err := neo4j.GetProperty[string](r, "uuid")
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
		Uuid:        UUID,
		Name:        name,
		Level:       level,
		Xp:          xp,
		Streak:      streak,
		LastDailyAt: lastDailyAt,
	}, nil
}

var DiscordQuery = &graphql.Field{
	Type: DiscordType,
	Args: graphql.FieldConfigArgument{
		"discordId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return repository.DiscordQuery(p), nil
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
		return repository.CreateDiscordMutation(p), nil
	},
}
