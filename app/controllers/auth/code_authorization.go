package auth

import (
	"context"
	"github.com/Nerzal/gocloak/v10"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"idist-go/app/controllers"
	"idist-go/app/providers/configProvider"
	"net/http"
)

type FormCodeAuthorization struct {
	Code string `json:"code" binding:"required"`
}

// CodeAuthorization godoc
// @Summary Lấy token từ code đo keycloak trả về
// @Accept  json
// @Produce  json
// @Router /api/v1/auth/code-authorization [post]
// @Success 200
//func CodeAuthorization(c *gin.Context) {
//	data := bson.M{}
//	config := configProvider.GetConfig()
//	entry := FormCodeAuthorization{}
//	// get code value
//	if err := c.ShouldBindBodyWith(&entry, binding.JSON); err != nil || entry.Code == "" {
//		controllers.ResponseError(c, http.StatusConflict, "Dữ liệu lỗi", err)
//	}
//	// Tạo request check
//	uri := fmt.Sprintf("%s/auth/realms/%s/protocol/openid-connect/token", config.GetString("keycloak.domain"), config.GetString("keycloak.realm"))
//	body := url.Values{}
//	body.Set("client_id", config.GetString("keycloak.client_id"))
//	body.Set("client_secret", config.GetString("keycloak.client_secret"))
//	body.Set("grant_type", "client_credentials")
//	body.Set("code", entry.Code)
//	body.Set("redirect_uri", config.GetString("keycloak.callback_url"))
//	body.Set("scope", "openid")
//
//	headers := map[string]string{}
//	response, err := helpers.HttpRequest(uri, http.MethodPost, headers, strings.NewReader(body.Encode()))
//	if err != nil {
//		controllers.ResponseError(c, http.StatusInternalServerError, "Gửi request lấy token lỗi", err)
//	}
//	f, err := ioutil.ReadAll(response.Body)
//
//	if response.StatusCode == 200 {
//		token := struct {
//			AccessToken      string `json:"access_token"`
//			RefreshToken     string `json:"refresh_token"`
//			Scope            string `json:"scope"`
//			IdToken          string `json:"id_token"`
//			TokenType        string `json:"token_type"`
//			ExpiresIn        int    `json:"expires_in"`
//			RefreshExpiresIn int    `json:"refresh_expires_in"`
//		}{}
//
//		_ = json.Unmarshal(f, &token)
//		data["access_token"] = token.AccessToken
//		data["refresh_token"] = token.RefreshToken
//		data["scope"] = token.Scope
//		data["id_token"] = token.IdToken
//		data["token_type"] = token.TokenType
//		data["expires_in"] = token.ExpiresIn
//		data["countdown"] = config.GetInt64("keycloak.countdown")
//
//		c.SetCookie("access_token", token.AccessToken, token.ExpiresIn, "/", c.Request.Host, true, false)
//		c.SetCookie("refresh_token", token.RefreshToken, token.RefreshExpiresIn, "/", c.Request.Host, true, false)
//		c.SetCookie("scope", token.Scope, 1000000, "/", c.Request.Host, true, false)
//		c.SetCookie("id_token", token.IdToken, 1000000, "/", c.Request.Host, true, false)
//		controllers.ResponseSuccess(c, http.StatusOK, "Đăng nhập thành công", data)
//	} else {
//		token := struct {
//			Error            string `json:"error"`
//			ErrorDescription string `json:"error_description"`
//		}{}
//		_ = json.Unmarshal(f, &token)
//		data["error"] = token.Error
//		data["error_description"] = token.ErrorDescription
//		controllers.ResponseError(c, response.StatusCode, "Đăng nhập thất bại", errors.New(token.ErrorDescription))
//	}
//
//}

func CodeAuthorization(c *gin.Context) {

	data := bson.M{}
	config := configProvider.GetConfig()
	client := gocloak.NewClient(config.GetString("keycloak.domain"))
	ctx := context.Background()
	token, err := client.Login(ctx, config.GetString("keycloak.client_id"), config.GetString("keycloak.client_secret"), config.GetString("keycloak.realm"), "toan.vuminh@savis.vn", "admin")
	if err != nil {
		panic("Login failed:" + err.Error())
	}

	tokenResult, err := client.RetrospectToken(ctx, token.AccessToken, config.GetString("keycloak.client_id"), config.GetString("keycloak.client_secret"), config.GetString("keycloak.realm"))

	if err != nil {
		controllers.ResponseError(c, http.StatusConflict, "Đăng nhập thất bại", err)
	}

	if !*tokenResult.Active {
		controllers.ResponseError(c, http.StatusConflict, "token is not active", err)
	}

	data["access_token"] = token.AccessToken
	data["refresh_token"] = token.RefreshToken
	data["scope"] = token.Scope
	data["id_token"] = token.IDToken
	data["token_type"] = token.TokenType
	data["expires_in"] = token.ExpiresIn
	data["countdown"] = config.GetInt64("keycloak.countdown")
	c.SetCookie("access_token", token.AccessToken, token.ExpiresIn, "/", c.Request.Host, true, false)
	c.SetCookie("refresh_token", token.RefreshToken, token.RefreshExpiresIn, "/", c.Request.Host, true, false)
	c.SetCookie("scope", token.Scope, 1000000, "/", c.Request.Host, true, false)
	c.SetCookie("id_token", token.IDToken, 1000000, "/", c.Request.Host, true, false)

	controllers.ResponseSuccess(c, http.StatusOK, "Đăng nhập thành công", data)
}
