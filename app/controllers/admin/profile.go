package admin

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"idist-go/app/models"
	"net/http"
)

func GetProfile(c *gin.Context) {
	data := bson.M{}
	user := c.MustGet("user").(models.User)
	data["entry"] = user
	ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}
