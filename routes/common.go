package routes

import (
	"github.com/gin-gonic/gin"
	"idist-core/app/controllers/admin"
	"idist-core/app/middlewares"
)

func CommonRoutes(router *gin.RouterGroup) {
	router.GET("/provinces", middlewares.Gate("", ""), admin.ListProvinces)
	router.GET("/provinces/id", middlewares.Gate("", ""), admin.ReadProvince)
	router.PUT("/provinces/id", middlewares.Gate("", ""), admin.UpdateProvince)

	router.POST("/tuyen-sinh", middlewares.Gate("", ""), admin.CreateTuyenSinh)
}
