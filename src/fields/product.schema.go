package fields

import "github.com/graphql-go/graphql"

var ProductType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Product",
	Fields: graphql.Fields{
		"id":              &graphql.Field{Type: graphql.String},
		"title":           &graphql.Field{Type: graphql.String},
		"description":     &graphql.Field{Type: graphql.String},
		"thumbnail":       &graphql.Field{Type: graphql.String},
		"price":           &graphql.Field{Type: graphql.String},
		"discountPercent": &graphql.Field{Type: graphql.String},
		"stock":           &graphql.Field{Type: graphql.String},
		"thubmnail":       &graphql.Field{Type: graphql.String},
		"status":          &graphql.Field{Type: graphql.String},
		"postion":         &graphql.Field{Type: graphql.String},
		"slug":            &graphql.Field{Type: graphql.String},
		"featured":        &graphql.Field{Type: graphql.String},
	},
})

var ProductInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "ProductInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"offset":   &graphql.InputObjectFieldConfig{Type: graphql.Int},
		"limit":    &graphql.InputObjectFieldConfig{Type: graphql.Int},
		"featured": &graphql.InputObjectFieldConfig{Type: graphql.String},
	},
})
