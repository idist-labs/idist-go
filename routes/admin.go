package routes

import (
	"ai-camera-api-cms/app/controllers/admin"
	"ai-camera-api-cms/app/middlewares"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.RouterGroup) {
	router.Use(middlewares.AuthMiddleware().MiddlewareFunc())
	router.GET("/refresh-token", middlewares.AuthMiddleware().RefreshHandler)
	router.GET("/profile", admin.GetProfile)
	router.POST("/logout", middlewares.AuthMiddleware().LogoutHandler)

	// Basic
	router.GET("/provinces", middlewares.Gate("", ""), admin.ListProvinces)
	router.GET("/provinces/id", middlewares.Gate("", ""), admin.ReadProvince)
	router.PUT("/provinces/id", middlewares.Gate("", ""), admin.UpdateProvince)

	router.GET("/districts", middlewares.Gate("", ""), admin.ListDistricts)
	router.GET("/districts/id", middlewares.Gate("", ""), admin.ReadDistrict)
	router.PUT("/districts/id", middlewares.Gate("", ""), admin.UpdateDistrict)

	router.GET("/wards", middlewares.Gate("", ""), admin.ListWards)
	router.GET("/wards/id", middlewares.Gate("", ""), admin.ReadWard)
	router.PUT("/wards/id", middlewares.Gate("", ""), admin.UpdateWard)

	router.GET("/roles", middlewares.Gate("", ""), admin.ListRoles)
	router.POST("/roles", middlewares.Gate("", ""), admin.CreateRole)
	router.GET("/roles/:id", middlewares.Gate("", ""), admin.ReadRole)
	router.PUT("/roles/:id", middlewares.Gate("", ""), admin.UpdateRole)
	router.DELETE("/roles/:id", middlewares.Gate("", ""), admin.DeleteRole)

	router.GET("/users", middlewares.Gate("", ""), admin.ListUsers)
	router.POST("/users", middlewares.Gate("", ""), admin.CreateUser)
	router.GET("/users/:id", middlewares.Gate("", ""), admin.ReadUser)
	router.PUT("/users/:id", middlewares.Gate("", ""), admin.UpdateUser)
	router.DELETE("/users/:id", middlewares.Gate("", ""), admin.DeleteUser)

	// Categories
	router.GET("/categories", middlewares.Gate("", ""), admin.ListCategories)
	router.POST("/categories", middlewares.Gate("", ""), admin.CreateCategory)
	router.GET("/categories/:id", middlewares.Gate("", ""), admin.GetCategory)
	router.PUT("/categories/:id", middlewares.Gate("", ""), admin.UpdateCategory)
	router.DELETE("/categories/:id", middlewares.Gate("", ""), admin.DeleteCategory)

	// Tags
	router.GET("/tags", middlewares.Gate("", ""), admin.ListTags)
	router.POST("/tags", middlewares.Gate("", ""), admin.CreateTag)
	router.GET("/tags/:id", middlewares.Gate("", ""), admin.GetTag)
	router.PUT("/tags/:id", middlewares.Gate("", ""), admin.UpdateTag)
	router.DELETE("/tags/:id", middlewares.Gate("", ""), admin.DeleteTag)

}
