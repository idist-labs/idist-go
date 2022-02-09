package admin

import (
	"ai-camera-api-cms/app/collections"
	"ai-camera-api-cms/app/controllers"
	"ai-camera-api-cms/const/mongo"
	"ai-camera-api-cms/const/response"
	"ai-camera-api-cms/helpers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strconv"
	"time"
)

func ListNotifications(c *gin.Context) {
	data := bson.M{}
	var err error
	entries := collections.Notifications{}
	entry := collections.Notification{}
	pagination := BindRequestTable(c, "created_at")
	filter := pagination.CustomFilters(bson.M{})
	opts := pagination.CustomOptions(options.Find())

	var listCameraIds []primitive.ObjectID
	checkListCameraId := true
	if c.Request.FormValue("province") != "" || c.Request.FormValue("district") != "" ||
		c.Request.FormValue("ward") != "" || c.Request.FormValue("village") != "" ||
		c.Request.FormValue("area") != "" {

		if listCameraIds, err = getListCameraId(c); err != nil {
			return
		}
		if len(listCameraIds) > 0 {
			filter["camera_id"] = bson.M{"$in": listCameraIds}
		} else {
			checkListCameraId = false
		}
	}
	if checkListCameraId == true {
		// search camera
		if c.Request.FormValue("camera") != "" {
			cameraId, err := primitive.ObjectIDFromHex(c.Request.FormValue("camera_id"))
			if err == nil {
				filter["camera_id"] = cameraId
			}
		}
		//Search time
		fromDate, _ := time.Parse("2006-01-02T15:04:05.000Z", c.Request.FormValue("from-date"))
		toDate, _ := time.Parse("2006-01-02T15:04:05.000Z", c.Request.FormValue("to-date"))
		if !toDate.IsZero() || !toDate.IsZero() {
			if toDate.IsZero() {
				toDate = helpers.Now()
			}
			filter["created_at"] = bson.M{
				"$gte": fromDate,
				"$lte": toDate,
			}
		}
		if entries, err = entry.Find(filter, opts); err != nil && err != mongo2.ErrNoDocuments {
			controllers.ResponseError(c, http.StatusInternalServerError, response.GET_FAIL, err)
			return
		}
	}
	pagination.Total, _ = entry.Count(filter)
	for i := 0; i < len(entries); i++ {
		entries[i].Preload("camera", "area")
	}
	data["entries"] = entries
	data["pagination"] = pagination
	controllers.ResponseSuccess(c, http.StatusOK, response.GET_SUCCESS, data)
}

func CreateNotification(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Notification{}

	if err = c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		controllers.ResponseError(c, http.StatusConflict, response.GET_FAIL, err)
	}

	// Xử lý tỉnh thành
	area := collections.Area{}
	filterArea := bson.M{
		"_id": entry.AreaId,
	}

	if err = area.First(filterArea); err != nil && err != mongo2.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusConflict, response.GET_FAIL, err)
	}
	entry.ProvinceId = area.ProvinceId
	entry.DistrictId = area.DistrictId
	entry.WardId = area.WardId
	entry.VillageId = area.VillageId

	if err = entry.Create(); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, response.CREATE_FAIL, err)
	}

	entry.Preload("area", "camera")
	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, response.CREATE_SUCCESS, data)
}
func ReadNotification(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Notification{}
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
	entry.Preload("camera")
	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, response.DELETE_SUCCESS, data)
}

func DeleteNotification(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Notification{}
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

	controllers.ResponseSuccess(c, http.StatusOK, response.DELETE_SUCCESS, data)
}

//func getListCameraId ()
func getListCameraId(c *gin.Context) (listCameraId []primitive.ObjectID, err error) {
	entry := collections.Camera{}
	entries := collections.Cameras{}
	filter := bson.M{
		"deleted_at": nil,
	}
	//Search
	if c.Request.FormValue("search") != "" {
		filter["name"] = bson.M{
			"$regex":   c.Request.FormValue("search"),
			"$options": "i",
		}
	}

	if c.Request.FormValue("province") != "" {
		provinceId, _ := strconv.ParseInt(c.Request.FormValue("province"), 10, 64)
		filter["province_id"] = provinceId
	}

	if c.Request.FormValue("district") != "" {
		districtId, _ := strconv.ParseInt(c.Request.FormValue("district"), 10, 64)
		filter["district_id"] = districtId
	}

	if c.Request.FormValue("ward") != "" {
		wardId, _ := strconv.ParseInt(c.Request.FormValue("ward"), 10, 64)
		filter["ward_id"] = wardId
	}

	if c.Request.FormValue("village") != "" {
		villageId, _ := strconv.ParseInt(c.Request.FormValue("village"), 10, 64)
		filter["village_id"] = villageId
	}
	fmt.Println(c.Request.FormValue("area"))
	if c.Request.FormValue("area") != "" {
		areaId, _ := primitive.ObjectIDFromHex(c.Request.FormValue("area"))
		filter["area_id"] = areaId
	}

	if entries, err = entry.Find(filter, nil); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, mongo.QUERY_FAIL, err)
	}

	if len(entries) > 0 {
		for i, _ := range entries {
			listCameraId = append(listCameraId, entries[i].ID)
		}
	}
	return listCameraId, err
}
