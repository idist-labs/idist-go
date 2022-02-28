package admin

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"idist-core/app/collections"
	"idist-core/app/controllers"
	"net/http"
)

func GetProfile(c *gin.Context) {
	fmt.Println("[GetProfile]")
	data := bson.M{}
	var err error
	entry := collections.User{}
	claims := jwt.ExtractClaims(c)
	userId, _ := primitive.ObjectIDFromHex(claims["id"].(string))

	filter := bson.M{
		"_id":        userId,
		"deleted_at": nil,
	}

	if err = entry.First(filter); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Trích xuất dữ liệu lỗi", nil)
	}
	data["entry"] = entry

	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}
