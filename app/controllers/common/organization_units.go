package common

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"gorm.io/gorm"
	"idist-go/app/controllers"
	"idist-go/app/models"
	"net/http"
)

// ListOrganizationUnits godoc
// @Summary Lấy danh sách đơn vị tổ chức
// @Accept  json
// @Produce  json
// @Router /api/v1/common/organization-units [get]
// @Success 200
// @TODO: Lấy danh sách đơn vị tổ chức
func ListOrganizations(c *gin.Context) {
	data := bson.M{}
	entry := models.Organization{}
	entries := models.Organizations{}
	var err error
	pagination := controllers.BindPagination(c)

	if entries, err = entry.Find(); err != nil && err != gorm.ErrRecordNotFound {
		controllers.ResponseError(c, http.StatusInternalServerError, "Lấy dữ liệu lỗi", err)
	}

	data["entries"] = entries
	data["pagination"] = pagination

	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}
