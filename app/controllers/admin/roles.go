package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"idist-core/app/collections"
	"idist-core/app/controllers"
	"idist-core/const/mongo"
	"idist-core/const/response"
	"net/http"
)

func ListRoles(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Role{}
	entries := collections.Roles{}
	filter := bson.M{
		"deleted_at": nil,
	}
	if entries, err = entry.Find(filter); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, mongo.QUERY_FAIL, err)
	}
	data["entries"] = entries
	controllers.ResponseSuccess(c, http.StatusOK, response.GET_SUCCESS, data)
}

func CreateRole(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Role{}

	if err = c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		controllers.ResponseError(c, http.StatusConflict, response.GET_FAIL, err)
	}

	if err = entry.Create(); err != nil {
		controllers.ResponseError(c, http.StatusConflict, response.CREATE_FAIL, err)
	}

	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, response.CREATE_SUCCESS, data)
}

func ReadRole(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Role{}
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

func UpdateRole(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Role{}
	exist := collections.Role{}
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

	if err = entry.Update(); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, response.UPDATE_FAIL, err)
	}

	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, response.UPDATE_SUCCESS, data)

}
func DeleteRole(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Role{}
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

	if err = entry.Delete(); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, response.DELETE_FAIL, err)
	}

	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, response.DELETE_SUCCESS, data)

}
