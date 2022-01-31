package routerProvider

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"idist-go/app/middlewares"
	"idist-go/app/providers/configProvider"
	"idist-go/routes"
)
import _ "idist-go/docs"

func Init(router *gin.Engine) {
	fmt.Println("------------------------------------------------------------")
	if configProvider.GetConfig().GetBool("app.server.log") {
		router.Use(gin.Logger())
	}
	router.Use(gin.Recovery())
	router.Use(middlewares.CorsMiddleware())
	router.Use(middlewares.ConfigsMiddleware())
	api := router.Group("api/v1")

	routes.AppRoutes(api.Group("app"))
	routes.AuthRoutes(api.Group("auth"))
	routes.CommonRoutes(api.Group("common"))
	routes.AdminRoutes(api.Group("admin"))

	// Swagger
	url := ginSwagger.URL("/api/swagger/doc.json") // The url pointing to API definition
	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	fmt.Println("------------------------------------------------------------")

}
