package routes

import (
	"ai-camera-api-cms/app/controllers/auth"
	"ai-camera-api-cms/app/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {
	router.POST("/register", auth.AuthRegister)
	router.POST("/login", middlewares.AuthMiddleware().LoginHandler)

}
