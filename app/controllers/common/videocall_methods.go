package common

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"idist-go/app/controllers"
	"net/http"
)

func ListVideoCallMethods(c *gin.Context) {
	data := bson.M{}
	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}
