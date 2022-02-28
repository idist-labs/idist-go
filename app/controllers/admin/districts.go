package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"idist-core/app/collections"
	"idist-core/app/controllers"
	"idist-core/const/mongo"
	"idist-core/const/response"
	"net/http"
	"strconv"
)

func ListDistricts(c *gin.Context) {
	data := bson.M{}
	var err error
	entries := collections.Districts{}
	entry := collections.District{}

	pagination := controllers.BindRequestTable(c, "_id")

	filter := pagination.CustomFilters(bson.M{})

	if c.Request.FormValue("province") != "" {
		provinceId, _ := strconv.Atoi(c.Request.FormValue("province"))
		filter["province_id"] = provinceId
	}
	if entries, err = entry.Find(filter); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, mongo.NOT_FOUND, err)
	}

	//for i := 0; i < len(entries); i++ {
	//	entries[i].Preload("wards")
	//}
	data["entries"] = entries
	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}

func ReadDistrict(c *gin.Context) {
	data := bson.M{}
	var err error
	entryId, _ := strconv.Atoi(c.Param("id"))
	entry := collections.District{}
	filter := bson.M{
		"_id": entryId,
	}

	if err = entry.First(filter); err != nil && err != mongo2.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, response.GET_FAIL, err)
	} else if err == mongo2.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusNotFound, mongo.NOT_FOUND, err)
	}

	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, response.GET_SUCCESS, data)
}
func UpdateDistrict(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.District{}
	exist := collections.District{}
	entryId, _ := strconv.Atoi(c.Param("id"))

	filter := bson.M{
		"_id":        entryId,
		"deleted_at": nil,
	}

	if err = exist.First(filter); err != nil && err != mongo2.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, response.GET_FAIL, err)
	} else if err == mongo2.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusNotFound, mongo.NOT_FOUND, err)
	}

	if err = c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		controllers.ResponseError(c, http.StatusConflict, response.GET_FAIL, err)
	}

	if err = entry.Update(); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, response.UPDATE_FAIL, err)
	}

	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, response.UPDATE_SUCCESS, data)

}
