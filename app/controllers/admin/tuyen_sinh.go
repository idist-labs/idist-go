package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	"idist-core/app/collections"
	"idist-core/app/controllers"
	"net/http"
)

func CreateTuyenSinh(c *gin.Context) {
	data := bson.M{}
	var err error

	entry := collections.TuyenSinh{}

	if err = c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		controllers.ResponseError(c, http.StatusBadRequest, "Dữ liệu gửi lên lỗi", err.Error())
	}

	if err = entry.Create(); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Xử lý dữ liệu lỗi", err)
	}

	data["entry"] = entry

	controllers.ResponseSuccess(c, http.StatusOK, "Đăng ký thành công", data)
}
