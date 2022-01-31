package routes

import (
	"github.com/gin-gonic/gin"
	"idist-go/app/controllers/admin"
	"idist-go/app/middlewares"
)

func AdminRoutes(router *gin.RouterGroup) {
	router.Use(middlewares.AuthorizationMiddleware())
	router.GET("/profile", admin.GetProfile)
}
