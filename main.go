package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"idist-go/app/providers/configProvider"
	"idist-go/app/providers/jobsProvider"
	"idist-go/app/providers/loggerProvider"
	"idist-go/app/providers/mysqlProvider"
	"idist-go/app/providers/postgresProvider"
	"idist-go/app/providers/redisProvider"
	"idist-go/app/providers/routerProvider"
	"idist-go/app/providers/socketProvider"
	"idist-go/database/migrations"
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
	loggerProvider.Init()
	// mongoProvider.Init()
	// defer mongoProvider.CloseMongoDB()
	mysqlProvider.Init()
	postgresProvider.Init()
	redisProvider.Init()
	jobsProvider.Init()
	routerProvider.Init(r)
	socketProvider.Init(r)

	if configProvider.GetConfig().GetBool("mysql.auto_migrate") {
		migrations.Runner()
	}
	///* Run server */
	if err := r.Run(fmt.Sprintf("%s:%s",
		configProvider.GetConfig().GetString("app.server.host"),
		configProvider.GetConfig().GetString("app.server.port"))); err != nil {
	}
}
