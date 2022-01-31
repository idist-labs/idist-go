package app

import (
	"github.com/Nerzal/gocloak/v10"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"idist-go/app/controllers"
	"net/http"
)

func GetProfile(c *gin.Context) {
	data := bson.M{}
	user := c.MustGet("user").(*gocloak.UserInfo)

	data["entry"] = user
	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}

func CreateScheduleProfile(c *gin.Context) {
	data := bson.M{}
	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}
