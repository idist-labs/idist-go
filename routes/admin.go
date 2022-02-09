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

	router.GET("/villages", middlewares.Gate("", ""), admin.ListVillages)
	router.POST("/villages", middlewares.Gate("", ""), admin.CreateVillage)
	router.GET("/villages/id", middlewares.Gate("", ""), admin.ReadVillage)
	router.PUT("/villages/id", middlewares.Gate("", ""), admin.UpdateVillage)

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

	router.GET("/areas", middlewares.Gate("", ""), admin.ListAreas)
	router.POST("/areas", middlewares.Gate("", ""), admin.CreateArea)
	router.GET("/areas/:id", middlewares.Gate("", ""), admin.ReadArea)
	router.PUT("/areas/:id", middlewares.Gate("", ""), admin.UpdateArea)
	router.DELETE("/areas/:id", middlewares.Gate("", ""), admin.DeleteArea)

	// Notifications
	router.GET("/notifications", middlewares.Gate("", ""), admin.ListNotifications)
	router.POST("/notifications", middlewares.Gate("", ""), admin.CreateNotification)
	router.GET("/notifications/:id", middlewares.Gate("", ""), admin.ReadNotification)
	router.DELETE("/notifications/:id", middlewares.Gate("", ""), admin.DeleteNotification)

	// Camera
	router.GET("/cameras", middlewares.Gate("", ""), admin.ListCameras)
	router.POST("/cameras", middlewares.Gate("", ""), admin.CreateCamera)
	router.GET("/cameras/:id", middlewares.Gate("", ""), admin.ReadCamera)
	router.PUT("/cameras/:id", middlewares.Gate("", ""), admin.UpdateCamera)
	router.DELETE("/cameras/:id", middlewares.Gate("", ""), admin.DeleteCamera)

	//API nhận thông tin Tracks và lấy link LIVE/PLAYBACK từ H-FACTOR
	router.GET("/tracks/live/:id", middlewares.Gate("", ""), admin.LiveTrack)
	router.GET("/tracks/view/:id/:time", middlewares.Gate("", ""), admin.ViewTrack)

	// Analytics
	router.GET("/securities/analytic-levels", middlewares.Gate("", ""), admin.AnalyticLevels)

}
