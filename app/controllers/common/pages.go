package common

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"gorm.io/gorm"
	response "idist-go/app/consts"
	"idist-go/app/controllers"
	"idist-go/app/models"
	"net/http"
)

func GetPage(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := models.Page{}
	if err = entry.First("slug", c.Param("slug")); err != nil && err != gorm.ErrRecordNotFound {
		controllers.ResponseError(c, http.StatusInternalServerError, response.GetError, err)
	} else if err == gorm.ErrRecordNotFound {
		controllers.ResponseError(c, http.StatusNotFound, response.GetError, nil)
	}

	data["entry"] = entry

	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}
