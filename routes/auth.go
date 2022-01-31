package routes

import (
	"github.com/gin-gonic/gin"
	"idist-go/app/controllers/auth"
)

func AuthRoutes(router *gin.RouterGroup) {
	router.POST("/login", auth.PostLogin)
	router.GET("/refresh-token", auth.GetRefreshToken)
	router.POST("/logout", auth.PostLogout)
}
