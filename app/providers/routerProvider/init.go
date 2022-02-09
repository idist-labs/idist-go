package routerProvider

import (
	"ai-camera-api-cms/app/middlewares"
	"ai-camera-api-cms/app/providers/configProvider"
	"ai-camera-api-cms/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	fmt.Println("------------------------------------------------------------")
	if configProvider.GetConfig().GetBool("app.server.log") {
		router.Use(gin.Logger())
	}
	router.Use(gin.Recovery())

	router.Use(middlewares.CorsMiddleware())
	router.Use(middlewares.ConfigsMiddleware())
	api := router.Group("api/v1")

	routes.AdminRoutes(api.Group("admin"))
	routes.AuthRoutes(api.Group("auth"))
	fmt.Println("------------------------------------------------------------")

}
