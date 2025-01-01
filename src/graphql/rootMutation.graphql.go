package graphql_config

import (
	"ThaiLy/resolvers"
	"ThaiLy/schemas"

	"github.com/graphql-go/graphql"
)

var rootMutation = graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"loginAccount": &graphql.Field{
			Type: schemas.AccountType,
			Args: graphql.FieldConfigArgument{
				"account": &graphql.ArgumentConfig{Type: graphql.NewNonNull(schemas.LoginAccountInput)},
			},
			Resolve: resolvers.LoginAccountResolver,
		},
		"registerAccount": &graphql.Field{
			Type: schemas.AccountType,
			Args: graphql.FieldConfigArgument{
				"account": &graphql.ArgumentConfig{Type: graphql.NewNonNull(schemas.RegisterAccountInput)},
			},
			Resolve: resolvers.RegisterAccountResolver,
		},
		"updateAccount": &graphql.Field{
			Type: schemas.AccountType,
			Args: graphql.FieldConfigArgument{
				"account": &graphql.ArgumentConfig{Type: schemas.UpdateAccountInput},
			},
			Resolve: resolvers.UpdateAccountResolver,
		},
		"createOtp": &graphql.Field{
			Type: schemas.OTPType,
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: resolvers.CreateOtpResolver,
		},
	},
}
