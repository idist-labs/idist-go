package main

import (
	"ai-camera-api-cms/app/providers/configProvider"
	"ai-camera-api-cms/app/providers/jobsProvider"
	"ai-camera-api-cms/app/providers/loggerProvider"
	"ai-camera-api-cms/app/providers/mongoProvider"
	"ai-camera-api-cms/app/providers/redisProvider"
	"ai-camera-api-cms/app/providers/routerProvider"
	"ai-camera-api-cms/app/providers/socketProvider"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	gin.DisableConsoleColor()
	if *environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	configProvider.Init(*environment)
	mongoProvider.Init()
	redisProvider.Init()
	jobsProvider.Init()
	loggerProvider.Init()
	routerProvider.Init(r)
	socketProvider.Init()

	defer mongoProvider.CloseMongoDB()
	///* Run server */
	if err := r.Run(fmt.Sprintf("%s:%s",
		configProvider.GetConfig().GetString("app.server.host"),
		configProvider.GetConfig().GetString("app.server.port"))); err != nil {
	}
}
