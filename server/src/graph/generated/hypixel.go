package generated

// Code automatically generated; DO NOT EDIT.

import (
	"github.com/graphql-go/graphql"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/src/db/repository"
)

type Player struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

var PlayerType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Player", Fields: graphql.Fields{
		"UUID": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
},
)

func ResultToPlayer(r *neo4j.EagerResult) (*Player, error) {
	result, _, err := neo4j.GetRecordValue[neo4j.Node](r.Records[0], "%!s(uint8=112)")
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

	return &Player{
		Uuid: UUID,
		Name: name,
	}, nil
}

var AccountMutation = &graphql.Field{
	Type: PlayerType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return repository.AccountMutation(p), nil
	},
}

var DiscordMutation = &graphql.Field{
	Type: PlayerType,
	Args: graphql.FieldConfigArgument{
		"discordId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return repository.DiscordMutation(p), nil
	},
}

var SigninMutation = &graphql.Field{
	Type: PlayerType,
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

var CreateDiscordMutation = &graphql.Field{
	Type: PlayerType,
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
