package auth

import (
	"ai-camera-api-cms/app/cache"
	"ai-camera-api-cms/app/collections"
	"ai-camera-api-cms/app/controllers"
	"ai-camera-api-cms/app/providers/configProvider"
	"errors"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

type Account struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Remember bool   `form:"remember" json:"remember"`
}

func PostLogin(c *gin.Context) (interface{}, error) {
	var account Account
	var err error
	entry := collections.User{}
	if err = c.ShouldBindBodyWith(&account, binding.JSON); err != nil {
		return "", err
	}

	filter := bson.M{
		"deleted_at": nil,
	}
	if strings.Contains(account.Username, "@") {
		filter["email"] = account.Username
	} else {
		filter["username"] = account.Username
	}
	// Truy vấn dữ liệu
	if err = entry.First(filter); err != nil && err != mongo.ErrNoDocuments {
		return nil, errors.New("Xử lý dữ liệu lỗi")
	} else if err == mongo.ErrNoDocuments {
		return nil, errors.New("Không tìm thấy tài khoản")
	}

	// Kiểm tra tài khoản
	if err = bcrypt.CompareHashAndPassword([]byte(entry.Password), []byte(account.Password)); err != nil {

		_ = entry.LoginFail()
		return nil, errors.New("Tài khoản hoặc mật khẩu không đúng")
	}

	// Kiểm tra điều kiện tài khoản
	if entry.Lock || entry.LoginFailedCount > configProvider.GetConfig().GetInt64("auth.login_fail") {
		return nil, errors.New("Tài khoản hiện đang bị khoá")
	}

	return entry, nil
}
func ResponseLogin(c *gin.Context, code int, token string, expire time.Time) {
	data := bson.M{
		"token":  token,
		"expire": expire.Format(time.RFC3339),
	}
	controllers.ResponseSuccess(c, http.StatusOK, "Đăng nhập thành công", data)

}
func Authorizator(data interface{}, c *gin.Context) bool {
	var err error
	redis := cache.GetInstance()
	cf := configProvider.GetConfig()
	userClaim, _ := data.(*collections.User)
	user := collections.User{}

	if err = redis.Get("user_"+userClaim.ID.String(), &user); err == nil {
		c.Set("user", user)
		_ = user.LastActive()
	} else {
		filter := bson.M{
			"_id":        userClaim.ID,
			"deleted_at": nil,
		}
		if err = user.First(filter); err != nil {
			fmt.Println(err)
			return false
		}
		user.Preload("permissions")
	}
	fmt.Println(user.Lock, user.LoginFailedCount > cf.GetInt64("auth.login_fail"))
	if user.Lock || user.LoginFailedCount > cf.GetInt64("auth.login_fail") {
		return false
	}
	c.Set("user", user)
	_ = redis.SetInterface("user_"+userClaim.ID.String(), user, 60)
	_ = user.LastActive()
	return true
}

func IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	entryId, _ := primitive.ObjectIDFromHex(claims["id"].(string))
	return &collections.User{
		ID:       entryId,
		Username: claims["username"].(string),
		Name:     claims["name"].(string),
	}
}

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(collections.User); ok {
		return jwt.MapClaims{
			"id":       v.ID,
			"username": v.Username,
			"name":     v.Name,
		}
	}
	return jwt.MapClaims{}
}

func Unauthorized(c *gin.Context, code int, message string) {
	controllers.ResponseError(c, code, message, nil)
}
