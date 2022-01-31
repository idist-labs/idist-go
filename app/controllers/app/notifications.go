package app

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"idist-go/app/controllers"
	"net/http"
)

func CheckNewNotifications(c *gin.Context) {
	data := bson.M{
		"has_new": 10,
	}
	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}
