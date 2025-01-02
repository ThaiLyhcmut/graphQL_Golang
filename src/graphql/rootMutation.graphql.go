package graphql_config

import (
	"ThaiLy/fields"
	"ThaiLy/resolvers"

	"github.com/graphql-go/graphql"
)

var rootMutation = graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"loginAccount": &graphql.Field{
			Type: fields.AccountType,
			Args: graphql.FieldConfigArgument{
				"account": &graphql.ArgumentConfig{Type: graphql.NewNonNull(fields.LoginAccountInput)},
			},
			Resolve: resolvers.LoginAccountResolver,
		},
		"registerAccount": &graphql.Field{
			Type: fields.AccountType,
			Args: graphql.FieldConfigArgument{
				"account": &graphql.ArgumentConfig{Type: graphql.NewNonNull(fields.RegisterAccountInput)},
			},
			Resolve: resolvers.RegisterAccountResolver,
		},
		"updateAccount": &graphql.Field{
			Type: fields.AccountType,
			Args: graphql.FieldConfigArgument{
				"account": &graphql.ArgumentConfig{Type: fields.UpdateAccountInput},
			},
			Resolve: resolvers.UpdateAccountResolver,
		},
		"createOtp": &graphql.Field{
			Type: fields.OTPType,
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: resolvers.CreateOtpResolver,
		},
	},
}
