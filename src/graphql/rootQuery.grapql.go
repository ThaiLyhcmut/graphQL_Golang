package graphql_config

import (
	"ThaiLy/fields"
	"ThaiLy/resolvers"

	"github.com/graphql-go/graphql"
)

var rootQuery = graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getAccount": &graphql.Field{
			Type:    fields.AccountType,
			Resolve: resolvers.GetAccountResolver,
		},
		"getCategory": &graphql.Field{
			Type: graphql.NewList(fields.CategoryType),
			Args: graphql.FieldConfigArgument{
				"categoryID": &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: resolvers.GetCategoryResolver,
		},
	},
}
