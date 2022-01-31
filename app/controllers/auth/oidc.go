package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"idist-go/app/controllers"
	"idist-go/app/providers/configProvider"
	"net/http"
)

// GetOidcURI godoc
// @Summary Lấy link OID
// @Accept  json
// @Produce  json
// @Router /api/v1/auth/oidc [get]
// @Success 200
func GetOidcURI(c *gin.Context) {
	config := configProvider.GetConfig()
	uri := fmt.Sprintf("%s/auth/realms/%s/protocol/openid-connect/auth?client_id=%s&redirect_uri=%s&response_type=code",
		config.GetString("keycloak.domain"),
		config.GetString("keycloak.realm"),
		config.GetString("keycloak.client_id"),
		config.GetString("keycloak.callback_url"))
	data := bson.M{
		"uri": uri,
	}
	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}
