package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"idist-core/app/providers/configProvider"
	"idist-core/app/providers/jobsProvider"
	"idist-core/app/providers/loggerProvider"
	"idist-core/app/providers/mongoProvider"
	"idist-core/app/providers/redisProvider"
	"idist-core/app/providers/routerProvider"
	"idist-core/app/providers/socketProvider"
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
