package routes

import (
	"github.com/gin-gonic/gin"
	"idist-go/app/controllers/common"
)

func CommonRoutes(router *gin.RouterGroup) {
	router.GET("/pages/:slug", common.GetPage)
	router.GET("/articles", common.ListArticles)
	router.GET("/articles/:id", common.GetArticle)
}
