package middlewares

import (
	"github.com/gin-gonic/gin"
)

func ConfigsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
