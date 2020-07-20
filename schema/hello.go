package schema

import (
	"github.com/graphql-go/graphql"
	"go-admin/models"
)

var helloType = graphql.NewObject(graphql.ObjectConfig{
	Name:       "Hello",
	Interfaces: nil,
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
	IsTypeOf:    nil,
	Description: "Hello Model",
})

var queryHello = graphql.Field{
	Name:        "QueryHello",
	Description: "Query Hello",
	Type:        graphql.NewList(helloType),
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		id, _ := p.Args["id"].(int)
		name, _ := p.Args["name"].(string)
		db := &models.Login{}
		db.Username = string(id)
		db.Password = name
		return nil, nil
	},
}
