package admin

import (
	"ai-camera-api-cms/app/collections"
	"ai-camera-api-cms/app/controllers"
	"ai-camera-api-cms/const/mongo"
	"ai-camera-api-cms/const/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"strings"
)

func ListUsers(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.User{}
	entries := collections.Users{}
	pagination := BindRequestTable(c, "created_at")
	filter := pagination.CustomFilters(bson.M{})
	opts := pagination.CustomOptions(options.Find())

	//Search
	if pagination.Search != "" {
		filter["name"] = bson.M{
			"$regex":   strings.TrimSpace(pagination.Search),
			"$options": "i",
		}
	}
	if c.Request.FormValue("province") != "" {
		provinceId, _ := strconv.Atoi(c.Request.FormValue("province"))
		filter["province_id"] = provinceId
	}

	if c.Request.FormValue("district") != "" {
		districtId, _ := strconv.Atoi(c.Request.FormValue("district"))
		filter["district_id"] = districtId
	}

	if c.Request.FormValue("ward") != "" {
		wardId, _ := strconv.Atoi(c.Request.FormValue("ward"))
		filter["ward_id"] = wardId
	}
	if c.Request.FormValue("village") != "" {
		villageId, _ := strconv.Atoi(c.Request.FormValue("village"))
		filter["village_id"] = villageId
	}
	if c.Request.FormValue("area") != "" {
		filter["area_id"] = strings.TrimSpace(c.Request.FormValue("area"))
	}

	if entries, err = entry.Find(filter, opts); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, mongo.QUERY_FAIL, err)
	}
	pagination.Total, _ = entry.Count(filter)
	for i := 0; i < len(entries); i++ {
		entries[i].Preload("province", "district", "ward", "village")
	}
	data["entries"] = entries
	data["pagination"] = pagination
	controllers.ResponseSuccess(c, http.StatusOK, response.GET_SUCCESS, data)
}

func CreateUser(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.User{}

	if err = c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		controllers.ResponseError(c, http.StatusConflict, response.GET_FAIL, err)
	}
	temp, _ := bcrypt.GenerateFromPassword([]byte(entry.NewPassword), bcrypt.DefaultCost)
	entry.Password = string(temp)

	// @TODO: Kiểm tra username có trùng ko, password có nhập ko

	if err = entry.Create(); err != nil {
		controllers.ResponseError(c, http.StatusConflict, response.CREATE_FAIL, err)
	}

	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, response.CREATE_SUCCESS, data)
}

func ReadUser(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.User{}
	entryId, _ := primitive.ObjectIDFromHex(c.Param("id"))

	filter := bson.M{
		"_id":        entryId,
		"deleted_at": nil,
	}

	if err = entry.First(filter); err != nil && err != mongo2.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, response.GET_FAIL, err)
	} else if err == mongo2.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusNotFound, mongo.NOT_FOUND, err)
	}

	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, response.GET_SUCCESS, data)
}

func UpdateUser(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.User{}
	exist := collections.User{}
	entryId, _ := primitive.ObjectIDFromHex(c.Param("id"))

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
	if entry.NewPassword != "" {
		temp, _ := bcrypt.GenerateFromPassword([]byte(entry.NewPassword), bcrypt.DefaultCost)
		entry.Password = string(temp)
	}

	// @TODO: Kiểm tra trùng username

	if err = entry.Update(); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, response.UPDATE_FAIL, err)
	}

	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, response.UPDATE_SUCCESS, data)

}
func DeleteUser(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.User{}
	entryId, _ := primitive.ObjectIDFromHex(c.Param("id"))
	user := c.MustGet("user").(collections.User)

	if user.ID == entryId {
		controllers.ResponseError(c, http.StatusConflict, "Không được tự xoá tài khoản đang đăng nhập", nil)
	}
	filter := bson.M{
		"_id":        entryId,
		"deleted_at": nil,
	}

	if err = entry.First(filter); err != nil && err != mongo2.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, response.GET_FAIL, err)
	} else if err == mongo2.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusNotFound, mongo.NOT_FOUND, err)
	}

	if err = entry.Delete(); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, response.DELETE_FAIL, err)
	}

	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, response.DELETE_SUCCESS, data)

}
