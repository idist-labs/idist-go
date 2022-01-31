package mysqlProvider

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
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
	fmt.Println("Mysql: Connecting...")
	c := configProvider.GetConfig()
	var err error
	CTimeOut = time.Duration(c.GetInt64("mysql.timeout")) * time.Second
	host := c.GetString("mysql.host")
	port := c.GetString("mysql.port")
	username := c.GetString("mysql.username")
	password := c.GetString("mysql.password")
	database := c.GetString("mysql.database")
	charset := c.GetString("mysql.charset")
	locale := c.GetString("mysql.locale")
	connect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s", username, password, host, port, database, charset, locale)

	logConf := logger.New(
		log.New(os.Stdout, "\r", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	if !c.GetBool("mysql.log") {
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
		mysql.Open(connect),
		&gorm.Config{
			Logger: logConf,
		},
	); err != nil {
		fmt.Println("Mysql: Connect to database fail.", zap.Error(err))
	} else {
		fmt.Println("Mysql: Connected!")
	}
}

func GetInstance() *gorm.DB {
	return instance
}
