package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"idist-core/app/collections"
	"idist-core/app/controllers"
	"idist-core/const/mongo"
	"idist-core/const/response"
	"net/http"
	"strconv"
)

func ListWards(c *gin.Context) {
	data := bson.M{}
	var err error
	entries := collections.Wards{}
	entry := collections.Ward{}

	pagination := controllers.BindRequestTable(c, "_id")

	filter := pagination.CustomFilters(bson.M{})
	opts := pagination.CustomOptions(options.Find())

	if c.Request.FormValue("district") != "" {
		districtId, _ := strconv.Atoi(c.Request.FormValue("district"))
		filter["district_id"] = districtId
	}
	if entries, err = entry.Find(filter, opts); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, mongo.NOT_FOUND, err)
	}

	data["entries"] = entries
	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}

func ReadWard(c *gin.Context) {
	data := bson.M{}
	var err error
	entryId, _ := strconv.Atoi(c.Param("id"))
	entry := collections.Ward{}
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
func UpdateWard(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Ward{}
	exist := collections.Ward{}
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
