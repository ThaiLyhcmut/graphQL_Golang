package fields

import (
	"ThaiLy/resolvers"

	"github.com/graphql-go/graphql"
)

var CategoryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Category",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.String},
		"title":       &graphql.Field{Type: graphql.String},
		"description": &graphql.Field{Type: graphql.String},
		"thumbnail":   &graphql.Field{Type: graphql.String},
		"status":      &graphql.Field{Type: graphql.String},
		"postion":     &graphql.Field{Type: graphql.String},
		"deleted":     &graphql.Field{Type: graphql.String},
		"slug":        &graphql.Field{Type: graphql.String},
		"product": &graphql.Field{
			Type: graphql.NewList(ProductType),
			Args: graphql.FieldConfigArgument{
				"ProductInput": &graphql.ArgumentConfig{Type: ProductInput},
			},
			Resolve: resolvers.GetProductByCategory,
		},
	},
})
