package graphql

import (
	"github.com/gin-gonic/gin"
	handlerx "github.com/graphql-go/handler"
	"go-admin/schema"
)

func GraphqlHandler() gin.HandlerFunc {
	h := handlerx.New(&handlerx.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
