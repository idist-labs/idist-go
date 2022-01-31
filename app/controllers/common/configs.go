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

func ListConfigs(c *gin.Context) {
	var err error
	entries := models.Configs{}
	entry := models.Config{}
	query := entry.Builder()

	query.Where("deleted_at IS NULL")
	query.Where("is_global = ?", true)

	if err = entry.FindGlobalConfigs(&entries); err != nil && err != gorm.ErrRecordNotFound {
		controllers.ResponseError(c, http.StatusInternalServerError, response.GetError, err)
	} else if err == gorm.ErrRecordNotFound {
		controllers.ResponseError(c, http.StatusNotFound, response.GetError, nil)
	}
	configs := bson.M{}
	for _, e := range entries {
		configs[e.Name] = e.Value
	}
	controllers.ResponseSuccess(c, http.StatusOK, response.GetSuccess, configs)
}
