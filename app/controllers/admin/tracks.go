package admin

import (
	"ai-camera-api-cms/app/collections"
	"ai-camera-api-cms/app/controllers"
	"ai-camera-api-cms/const/response"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

//Hàm lấy thông tin Track từ id (startTime == "" => LIVE ngược lại là PLAYBACK)
//startTime có dạng yyyyMMddThhmmssZ ví dụ 20211107T091530Z => 09 giờ 15 phút 30 giây ngày 07 tháng 10 năm 2021

func LiveTrack(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Camera{}
	entryId, _ := primitive.ObjectIDFromHex(c.Param("id"))

	filter := bson.M{
		"_id":        entryId,
		"deleted_at": nil,
	}
	if err = entry.First(filter); err != nil && err != mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, response.GET_FAIL, err)
	} else if err == mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusNotFound, response.GET_FAIL, err)
	}

	if link := entry.RenderStreamUrl(""); link == "" {
		controllers.ResponseError(c, http.StatusNotFound, "Không tìm thấy luổng video", nil)
	} else {
		data["link"] = link
		controllers.ResponseSuccess(c, http.StatusOK, response.GET_SUCCESS, data)
	}
}

func ViewTrack(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Camera{}
	entryId, _ := primitive.ObjectIDFromHex(c.Param("id"))
	time := c.Param("time")
	filter := bson.M{
		"_id":        entryId,
		"deleted_at": nil,
	}
	if err = entry.First(filter); err != nil && err != mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, response.GET_FAIL, err)
	} else if err == mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusNotFound, response.GET_FAIL, err)
	}

	if link := entry.RenderStreamUrl(time); link == "" {
		controllers.ResponseError(c, http.StatusNotFound, "Không tìm thấy luổng video", nil)
	} else {
		data["link"] = link
		controllers.ResponseSuccess(c, http.StatusOK, response.GET_SUCCESS, data)
	}
}
