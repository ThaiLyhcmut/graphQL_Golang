package fields

import "github.com/graphql-go/graphql"

var AccountType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Account",
	Fields: graphql.Fields{
		"id":       &graphql.Field{Type: graphql.String},
		"fullName": &graphql.Field{Type: graphql.String},
		"email":    &graphql.Field{Type: graphql.String},
		"adress":   &graphql.Field{Type: graphql.String},
		"phone":    &graphql.Field{Type: graphql.String},
		"avatar":   &graphql.Field{Type: graphql.String},
		"sex":      &graphql.Field{Type: graphql.String},
		"birthday": &graphql.Field{Type: graphql.String},
		"token":    &graphql.Field{Type: graphql.String},
		"code":     &graphql.Field{Type: graphql.String},
		"msg":      &graphql.Field{Type: graphql.String},
	},
})

var RegisterAccountInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "RegisterAccountInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"fullName": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"email":    &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"password": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"otp":      &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
	},
})

var LoginAccountInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "LoginAccountInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"email":    &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"password": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
	},
})

var UpdateAccountInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UpdateAccountInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"fullName": &graphql.InputObjectFieldConfig{Type: graphql.String},
		"adress":   &graphql.InputObjectFieldConfig{Type: graphql.String},
		"phone":    &graphql.InputObjectFieldConfig{Type: graphql.String},
		"avatar":   &graphql.InputObjectFieldConfig{Type: graphql.String},
		"sex":      &graphql.InputObjectFieldConfig{Type: graphql.String},
		"birthday": &graphql.InputObjectFieldConfig{Type: graphql.String},
	},
})

var OTPType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OTP",
	Fields: graphql.Fields{
		"code": &graphql.Field{Type: graphql.String},
		"msg":  &graphql.Field{Type: graphql.String},
	},
})

var CreateOtpInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CreateOtpInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"email": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
	},
})
