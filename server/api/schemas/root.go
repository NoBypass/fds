package schemas

import "github.com/graphql-go/graphql"

var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"hello": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "Hello, GraphQL!", nil
				},
			},
			"account": AccountQueryByUsername,
			"discord": DiscordQueryByDiscordId,
		},
	},
)

var rootMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createMessage": &graphql.Field{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"content": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					content, _ := p.Args["content"].(string)
					return "New message created: " + content, nil
				},
			},
			"createAccount": CreateAccount,
			"createDiscord": CreateDiscord,
		},
	},
)

var RootSchema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	},
)
