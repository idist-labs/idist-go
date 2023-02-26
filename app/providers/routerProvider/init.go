package routerProvider

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"idist-core/app/middlewares"
	"idist-core/app/providers/configProvider"
	_ "idist-core/docs"
	"idist-core/routes"
)

func Init(router *gin.Engine) {
	fmt.Println("------------------------------------------------------------")
	if configProvider.GetConfig().GetBool("app.server.log") {
		router.Use(gin.Logger())
	}
	router.Use(gin.Recovery())

	router.Use(middlewares.CorsMiddleware())
	router.Use(middlewares.ConfigsMiddleware())
	//docs.SwaggerInfo.BasePath = "swagger/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := router.Group("api/v1")
	routes.AdminRoutes(api.Group("admin"))
	routes.AuthRoutes(api.Group("auth"))
	routes.CommonRoutes(api.Group("common"))
	fmt.Println("------------------------------------------------------------")

}
