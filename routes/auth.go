package routes

import (
	"github.com/gin-gonic/gin"
	"idist-core/app/controllers/auth"
	"idist-core/app/middlewares"
)

func AuthRoutes(router *gin.RouterGroup) {
	router.POST("/register", auth.AuthRegister)
	router.POST("/login", middlewares.AuthMiddleware().LoginHandler)

}
