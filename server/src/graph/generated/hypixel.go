package generated

// Code automatically generated; DO NOT EDIT.

import (
	"github.com/graphql-go/graphql"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/src/repository"
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

func ResultToPlayer(result *neo4j.EagerResult) (*Player, error) {
	r, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "p")
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

	return &Player{
		Uuid: UUID,
		Name: name,
	}, nil
}

type PlayerInput struct {
	Name string `json:"name"`
}

var PlayerQuery = &graphql.Field{
	Type: PlayerType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		input := &PlayerInput{
			Name: p.Args["name"].(string)}

		return repository.PlayerQuery(&p.Context, input), nil
	},
}
