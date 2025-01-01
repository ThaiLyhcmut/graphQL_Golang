package graphql_config

import (
	"github.com/graphql-go/graphql"
)

func Config() *graphql.Schema {
	var config, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    graphql.NewObject(rootQuery),
			Mutation: graphql.NewObject(rootMutation),
		},
	)
	return &config
}
