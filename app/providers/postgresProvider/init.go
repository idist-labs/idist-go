package postgresProvider

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"idist-go/app/providers/configProvider"
	"log"
	"os"
	"time"
)

var instance *gorm.DB

var CTimeOut = 10 * time.Second

func Init() {
	//ctx, cancel := context.WithTimeout(context.Background(), CTimeOut)
	//defer cancel()
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Postgres: Connecting...")
	c := configProvider.GetConfig()
	var err error
	CTimeOut = time.Duration(c.GetInt64("postgres.timeout")) * time.Second
	host := c.GetString("postgres.host")
	port := c.GetString("postgres.port")
	username := c.GetString("postgres.username")
	password := c.GetString("postgres.password")
	database := c.GetString("postgres.database")
	locale := c.GetString("postgres.locale")
	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", host, port, username, password, database, locale)
	logConf := logger.New(
		log.New(os.Stdout, "\r", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	if !c.GetBool("postgres.log") {
		logConf = logger.New(
			log.New(os.Stdout, "\r", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Warn, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			},
		)
	}

	if instance, err = gorm.Open(
		postgres.Open(connect),
		&gorm.Config{
			Logger: logConf,
		},
	); err != nil {
		fmt.Println("Postgres: Connect to database fail.", zap.Error(err))
	} else {
		fmt.Println("Postgres: Connected!")
	}
}

func GetInstance() *gorm.DB {
	return instance
}
