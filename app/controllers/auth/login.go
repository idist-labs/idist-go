package auth

import (
	"context"
	"github.com/Nerzal/gocloak/v10"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	"gorm.io/gorm"
	"idist-go/app/controllers"
	"idist-go/app/models"
	"idist-go/app/providers/configProvider"
	"idist-go/helpers"
	"net/http"
)

type FormAccount struct {
	Username string `json:"username"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Remember bool   `json:"remember"`
}

func PostLogin(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := models.User{}
	now := helpers.Now()
	account := FormAccount{}

	if err = c.ShouldBindBodyWith(&account, binding.JSON); err != nil {
		controllers.ResponseError(c, http.StatusConflict, "Tài khoản không đúng", err)
	}

	// Gọi Keycloak
	config := configProvider.GetConfig()
	keycloak := gocloak.NewClient(config.GetString("keycloak.domain"))
	ctx := context.Background()
	token, err := keycloak.Login(ctx,
		config.GetString("keycloak.client_id"),
		config.GetString("keycloak.client_secret"),
		config.GetString("keycloak.realm"),
		account.Username,
		account.Password)

	if err != nil {
		controllers.ResponseError(c, http.StatusUnauthorized, "Đăng nhập thất bại", err)
	}

	// Check thông tin người dùng
	user, err := keycloak.GetUserInfo(c, token.AccessToken, config.GetString("keycloak.realm"))
	if err = entry.First("sub", *user.Sub); err != nil && err != gorm.ErrRecordNotFound {
		controllers.ResponseError(c, http.StatusUnauthorized, "Kiểm tra thông tin thất bại", err)
	} else if err == gorm.ErrRecordNotFound {
		entry = models.User{
			Sub:                 gocloak.PString(user.Sub),
			GivenName:           gocloak.PString(user.GivenName),
			FamilyName:          gocloak.PString(user.FamilyName),
			MiddleName:          gocloak.PString(user.MiddleName),
			Nickname:            gocloak.PString(user.Nickname),
			PreferredUsername:   gocloak.PString(user.PreferredUsername),
			Profile:             gocloak.PString(user.Profile),
			Picture:             gocloak.PString(user.Picture),
			Website:             gocloak.PString(user.Website),
			Email:               gocloak.PString(user.Email),
			EmailVerified:       gocloak.PBool(user.EmailVerified),
			Gender:              gocloak.PString(user.Gender),
			ZoneInfo:            gocloak.PString(user.ZoneInfo),
			Locale:              gocloak.PString(user.Locale),
			PhoneNumber:         gocloak.PString(user.PhoneNumber),
			PhoneNumberVerified: gocloak.PBool(user.PhoneNumberVerified),
			Name:                gocloak.PString(user.Name),
			Username:            gocloak.PString(user.PreferredUsername),
			LastActiveAt:        &now,
			CreatedAt:           helpers.Now(),
			UpdatedAt:           helpers.Now(),
		}
		if err = entry.Insert(); err != nil {
			controllers.ResponseError(c, http.StatusUnauthorized, "Đồng bộ thông tin thất bại", err)
		}
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
