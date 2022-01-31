package admin

import (
	"github.com/gin-gonic/gin"
	"idist-go/app/providers/configProvider"
)

func ResponseSuccess(c *gin.Context, code int, message string, data interface{}) {
	con := configProvider.GetConfig()
	c.AbortWithStatusJSON(code, gin.H{
		"code":    code,
		"data":    data,
		"message": message,
		"version": con.GetString("app.server.version"),
	})
	panic(nil)
}

func ResponseError(c *gin.Context, code int, message string, error interface{}) {
	con := configProvider.GetConfig()
	if con.GetString("app.env") == "production" {
		error = nil
	}
	c.AbortWithStatusJSON(code, gin.H{
		"code":    code,
		"error":   error,
		"message": message,
		"version": con.GetString("app.server.version"),
	})
	panic(nil)
}
