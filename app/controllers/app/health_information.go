package app

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	"idist-go/app/controllers"
	"idist-go/app/models"
	"idist-go/app/providers/configProvider"
	"idist-go/helpers"
	"io/ioutil"
	"net/http"
)

func GetHealthInformation(c *gin.Context) {
	data := bson.M{}
	config := configProvider.GetConfig()

	user := c.MustGet("user").(models.User)
	// @TODO: lấy thông tin
	body, _ := json.Marshal(user)

	headers := map[string]string{"Content-Type": "application/json"}
	response, err := helpers.HttpRequest(config.GetString("cert.signing"), http.MethodPost, headers, bytes.NewBuffer(body))
	if err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Gửi request lấy token lỗi", err)
	}
	f, err := ioutil.ReadAll(response.Body)
	if response.StatusCode == 200 {
		data["value"] = string(f)
	} else {
		controllers.ResponseError(c, http.StatusInternalServerError, "Mã hoá lỗi", err)
	}
	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}

func DecodeHealthInformation(c *gin.Context) {
	config := configProvider.GetConfig()
	//bind data
	healthInfo := struct {
		Value string `json:"value" binding:"required"`
	}{}

	if err := c.ShouldBindBodyWith(&healthInfo, binding.JSON); err != nil {
		controllers.ResponseError(c, http.StatusConflict, "Dữ liệu gửi lên không đúng", err)
	}

	body := []byte(`{"data":"` + healthInfo.Value + `"}`)

	headers := map[string]string{"Content-Type": "application/json"}
	response, err := helpers.HttpRequest(config.GetString("cert.verification"), http.MethodPost, headers, bytes.NewBuffer(body))
	if err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Gửi request giải mã lỗi", err)
	}
	f, err := ioutil.ReadAll(response.Body)
	if err == nil && response.StatusCode == 200 {
		var tmpe interface{}
		_ = json.Unmarshal(f, &tmpe)
		controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", tmpe)

	} else {
		controllers.ResponseError(c, http.StatusInternalServerError, "Giải hoá lỗi", err)
	}
}
