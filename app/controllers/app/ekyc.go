package app

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"idist-go/app/controllers"
	"mime/multipart"
	"net/http"
)

type FormEkyc struct {
	ImageCard1   *multipart.FileHeader `form:"image_card1" binding:"required"`
	ImageCard2   *multipart.FileHeader `form:"image_card2" binding:"required"`
	ImageGeneral *multipart.FileHeader `form:"image_general" binding:"required"`
}

func PostVerifyKYC(c *gin.Context) {
	data := bson.M{}
	formEkyc := FormEkyc{}
	var err error

	if formEkyc.ImageCard1, err = c.FormFile("image_card1"); err != nil {

	}
	if formEkyc.ImageCard1, err = c.FormFile("image_card2"); err != nil {

	}
	if formEkyc.ImageCard1, err = c.FormFile("image_general"); err != nil {

	}

	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}

func UpdateVerifyKYC(c *gin.Context) {
	data := bson.M{}
	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}
