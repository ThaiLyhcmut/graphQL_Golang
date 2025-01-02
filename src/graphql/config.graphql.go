package graphql_config

import (
	"github.com/graphql-go/graphql"
)

func Config() *graphql.Schema {
	// tạo 1 Schema mới
	var config, _ = graphql.NewSchema(
		// Tạo 1 Schema Config gồm các thành phần của Schema Như Query, Mutaion
		graphql.SchemaConfig{
			// Tạo thành
			Query:    graphql.NewObject(rootQuery), // Mỗi cái này là các Object mới
			Mutation: graphql.NewObject(rootMutation),
		},
	)
	return &config
}
