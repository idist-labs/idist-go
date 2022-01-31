package app

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"idist-go/app/controllers"
	"net/http"
)

type AnalyticMedicalRecords struct {
	Total int64 `json:"total"`
}

func GetAnalyticMedicalRecords(c *gin.Context) {
	data := bson.M{}
	data["entry"] = AnalyticMedicalRecords{}
	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}
