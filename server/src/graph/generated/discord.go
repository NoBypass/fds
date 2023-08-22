package generated

// Code automatically generated; DO NOT EDIT.

import (
	"github.com/graphql-go/graphql"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Discord struct {
	Uuid        string `json:"uuid"`
	Name        string `json:"name"`
	Level       int64  `json:"level"`
	Xp          int64  `json:"xp"`
	Streak      int64  `json:"streak"`
	LastDailyAt string `json:"last_daily_at"`
}

var discordType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Discord",
		Fields: graphql.Fields{
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
	accountNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "a")
	if err != nil {
		return nil, err
	}

	UUID, err := neo4j.GetProperty[string](accountNode, "uuid")
	if err != nil {
		return nil, err
	}

	name, err := neo4j.GetProperty[string](accountNode, "name")
	if err != nil {
		return nil, err
	}

	level, err := neo4j.GetProperty[int64](accountNode, "level")
	if err != nil {
		return nil, err
	}

	xp, err := neo4j.GetProperty[int64](accountNode, "xp")
	if err != nil {
		return nil, err
	}

	streak, err := neo4j.GetProperty[int64](accountNode, "streak")
	if err != nil {
		return nil, err
	}

	lastDailyAt, err := neo4j.GetProperty[string](accountNode, "last_daily_at")
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
