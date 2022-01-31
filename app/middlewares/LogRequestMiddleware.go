package middlewares

import (
	"github.com/gin-gonic/gin"
	"idist-go/app/providers/configProvider"
	"strings"
)

func getClientIp(c *gin.Context) string {
	if c.Request.Header.Get("CF-Connecting-IP") != "" {
		return c.Request.Header.Get("CF-Connecting-IP")
	}
	if c.Request.Header.Get("X-Forwarded-For") != "" {
		return c.Request.Header.Get("X-Forwarded-For")
	}
	if c.Request.Header.Get("X-Real-Ip") != "" {
		return c.Request.Header.Get("X-Real-Ip")
	}
	s := strings.Split(c.Request.RemoteAddr, ":")
	return s[0]
}
func LogRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" && !configProvider.GetConfig().GetBool("log.methods.get") {
			c.Next()
			return
		}
		//go func() {
		//	//DB := database.GetMongoDB()
		//	// Save log request
		//	log := collections.SystemLog{}
		//
		//	log.Method = c.Request.Method
		//	log.Api = c.Request.RequestURI
		//	log.ConvertName()
		//	log.IP = getClientIp(c)
		//	log.UserAgent = c.Request.UserAgent()
		//
		//	// Xử lý param request
		//	log.Params = c.Request.URL.Query().Encode()
		//
		//	//// Lưu vào DB
		//	//if err := log.Create(DB); err != nil {
		//	//	utils.Logger.Error("SaveRequest: lưu log không admin lỗi", zap.Error(err))
		//	//}
		//}()

		c.Next()
	}
}

//func saveFile(logS collections.SystemLog, admin collections.Account) string {
//
//	if dir, err := os.Getwd(); err != nil {
//		loggerProvider.Logger.Error("SaveRequest: get current dir error", zap.Error(err))
//		return ""
//	} else {
//		currentTime := time.Now().Format("2006-01-02 15:04:05.999999999")
//		currentDir := filepath.Join(dir, "public", "history", admin.ID.Hex())
//		// Tạo thư mục
//		_ = os.MkdirAll(currentDir, 0775)
//
//		logHistory := map[string]interface{}{
//			"data": logS.Query,
//			"user": admin,
//			"api":  logS.Api,
//		}
//		if logHistoryString, err := json.Marshal(logHistory); err != nil {
//			//utils.Logger.Error("SaveRequest: run marshal error", zap.Error(err))
//			return ""
//		} else {
//			if err := ioutil.WriteFile(filepath.Join(currentDir, currentTime+".json"), logHistoryString, 0644); err != nil {
//				//utils.Logger.Error("SaveRequest: write data error", zap.Error(err))
//				return ""
//			} else {
//				return filepath.Join(currentDir, currentTime+".json")
//			}
//		}
//	}
//}
