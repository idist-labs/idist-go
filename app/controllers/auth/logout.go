package auth

import (
	"github.com/Nerzal/gocloak/v10"
	"github.com/gin-gonic/gin"
	"idist-go/app/controllers"
	"idist-go/app/providers/configProvider"
	"net/http"
)

func PostLogout(c *gin.Context) {
	config := configProvider.GetConfig()
	client := gocloak.NewClient(config.GetString("keycloak.domain"))
	refreshToken, _ := c.Cookie("refresh_token")

	if refreshToken == "" {
		refreshToken = c.Request.Header.Get("RefreshToken")
	}

	err := client.Logout(c, config.GetString("keycloak.client_id"), config.GetString("keycloak.client_secret"), config.GetString("keycloak.realm"), refreshToken)
	// Xoá Token

	if err != nil {
		controllers.ResponseError(c, http.StatusBadRequest, "Đăng xuất không thành công", err)
	}
	c.SetCookie("access_token", "", -1, "/", c.Request.Host, true, false)
	c.SetCookie("refresh_token", "", -1, "/", c.Request.Host, true, false)
	c.SetCookie("scope", "", -1, "/", c.Request.Host, true, false)
	c.SetCookie("id_token", "", -1, "/", c.Request.Host, true, false)

	controllers.ResponseSuccess(c, http.StatusOK, "Đăng xuất thành công", nil)
}
