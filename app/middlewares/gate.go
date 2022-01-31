package middlewares

import (
	"github.com/gin-gonic/gin"
	"idist-go/app/models"
)

func Gate(Subject, Action string) gin.HandlerFunc {
	user := models.User{}
	return func(c *gin.Context) {
		if Subject == "" || Action == "" {
			c.Next()
			return
		}
		if _, exist := c.Get("user"); exist == true {
			user = c.MustGet("user").(models.User)
		}
		//if user.HasPermission(models.Permission{Subject: Subject, Action: Action}) {
		c.Next()
		return
		//}
		c.AbortWithStatus(403)
		return
	}
}
