package app

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"idist-go/app/controllers"
	"net/http"
)

func ListCertificates(c *gin.Context) {
	data := bson.M{}
	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}

func ListCertificateForms(c *gin.Context) {
	data := bson.M{}
	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}

func CreateCertificate(c *gin.Context) {
	data := bson.M{}
	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}

func ReadCertificate(c *gin.Context) {
	data := bson.M{}
	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}

func UpdateCertificate(c *gin.Context) {
	data := bson.M{}
	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}

func DeleteCertificate(c *gin.Context) {
	data := bson.M{}
	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}
