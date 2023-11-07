package configs

import (
	"github.com/gin-gonic/gin"
	"github.com/oat431/api-tutorial-go/schema"
	"github.com/graphql-go/handler"
)

func SetupGql() gin.HandlerFunc {
	return func(c *gin.Context) {
			h := handler.New(&handler.Config{
				Schema:   schema.CreateRootSchema().Root(),
				Pretty:   true,
				GraphiQL: true,
			})
			h.ServeHTTP(c.Writer, c.Request)
		}
}