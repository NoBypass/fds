package schemas

import (
	"github.com/graphql-go/graphql"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/db/repository"
)

var playerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Player",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"uuid": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var PlayerQueryByName = &graphql.Field{
	Type: playerType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return repository.FindPlayerByName(p.Context, p.Context.Value("driver").(neo4j.DriverWithContext), p.Args["name"].(string))
	},
}
