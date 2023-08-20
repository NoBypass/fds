package generated

// Code automatically generated; DO NOT EDIT.

import (
	"github.com/graphql-go/graphql"
	"server/src/db/repository"
)

type Player struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}


var playerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Player",
		Fields: graphql.Fields{
			"uuid": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
)


var PlayerQuery = &graphql.Field{
	Type: playerType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return repository.PlayerQuery(p)
	},
}


func ResultToPlayer(result *neo4j.EagerResult) (*Player, error) {	accountNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "a")
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

	return &Player{
		Uuid: UUID,
		Name: name,
	}, nil
}


