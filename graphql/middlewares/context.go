package middlewares

import (
	"github.com/gin-gonic/gin"
)

// GinContextToContext is the middleware to convert gin context to go context
func GinContextToContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := utils.CreateContextFromGinContext(c)
		ctx = utils.WithRequest(ctx, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
