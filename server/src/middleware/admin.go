package middleware

import (
	"github.com/gin-gonic/gin"
	"server/src/models"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		value, exists := c.Get("user")
		if exists && value.(models.User).Admin {
			c.Next()
		}
		c.Abort()
	}
}
