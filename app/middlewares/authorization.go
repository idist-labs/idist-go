package middlewares

import (
	"github.com/Nerzal/gocloak/v10"
	"github.com/gin-gonic/gin"
	"idist-go/app/models"
	"idist-go/app/providers/configProvider"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		config := configProvider.GetConfig()
		client := gocloak.NewClient(config.GetString("keycloak.domain"))
		accessToken, _ := c.Cookie("access_token")
		if accessToken == "" {
			if c.Request.Header.Get("Authorization") == "" {
				c.AbortWithStatus(401)
				return
			}
			accessToken = c.Request.Header.Get("Authorization")[6:]
		}
		user, err := client.GetUserInfo(c, accessToken, config.GetString("keycloak.realm"))
		if err != nil {
			c.AbortWithStatus(401)
		}
		err = VerifyUser(c, user)

		if err != nil {
			c.AbortWithStatus(401)
		}
		c.Next()
	}
}

func VerifyUser(c *gin.Context, user *gocloak.UserInfo) error {
	// @TODO Kiểm tra user có tồn tại thì load thông tin ra, nếu chưa tồn tại thì lưu vào DB
	entry := models.User{}

	if err := entry.First("sub", *user.Sub); err != nil {
		c.AbortWithStatus(500)
		return err
	}

	c.Set("user", entry)
	return nil
}
