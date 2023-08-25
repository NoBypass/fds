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

func ResultToDiscord(r *neo4j.EagerResult) (*Discord, error) {
	result, _, err := neo4j.GetRecordValue[neo4j.Node](r.Records[0], "%!s(uint8=100)")
	if err != nil {
		return nil, err
	}

	UUID, err := neo4j.GetProperty[string](result, "uuid")
	if err != nil {
		return nil, err
	}

	name, err := neo4j.GetProperty[string](result, "name")
	if err != nil {
		return nil, err
	}

	level, err := neo4j.GetProperty[int64](result, "level")
	if err != nil {
		return nil, err
	}

	xp, err := neo4j.GetProperty[int64](result, "xp")
	if err != nil {
		return nil, err
	}

	streak, err := neo4j.GetProperty[int64](result, "streak")
	if err != nil {
		return nil, err
	}

	lastDailyAt, err := neo4j.GetProperty[string](result, "last_daily_at")
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

var AccountMutation = &graphql.Field{
	Type: DiscordType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return repository.AccountMutation(p), nil
	},
}

var SigninMutation = &graphql.Field{
	Type: DiscordType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"remember": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return repository.SigninMutation(p), nil
	},
}
