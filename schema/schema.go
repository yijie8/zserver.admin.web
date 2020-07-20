package schema

import "github.com/graphql-go/graphql"

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:       "RootQuery",
	Interfaces: nil,
	Fields: graphql.Fields{
		"hello": &queryHello,
	},
	IsTypeOf:    nil,
	Description: "Root Query",
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: nil,
})
