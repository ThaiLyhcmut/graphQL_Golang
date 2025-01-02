package graphql_config

import (
	"ThaiLy/resolvers"
	"ThaiLy/schemas"

	"github.com/graphql-go/graphql"
)

var rootQuery = graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getAccount": &graphql.Field{
			Type:    schemas.AccountType,
			Resolve: resolvers.GetAccountResolver,
		},
		"getCategory": &graphql.Field{
			Type: graphql.NewList(schemas.CategoryType),
			Args: graphql.FieldConfigArgument{
				"categoryID": &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: resolvers.GetCategoryResolver,
		},
	},
}
