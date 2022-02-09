package configProvider

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var config *viper.Viper

func Init(env string) {
	fmt.Println("------------------------------------------------------------")
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	files, err := ioutil.ReadDir(getConfigDir(env))
	if err != nil {
		log.Fatal(err)
	}
	config.SetConfigName("app")
	config.AddConfigPath("configs/" + env + "/")
	_ = config.ReadInConfig()

	for _, f := range files {
		config.SetConfigName(f.Name())
		_ = config.MergeInConfig()
		fmt.Println("Loaded config: " + f.Name())
	}

	if err != nil {
		log.Fatal("error on parsing configuration file", zap.Error(err))
	}

}

func getConfigDir(env string) string {
	configDir := ""
	if dir, err := os.Getwd(); err != nil {
		return ""
	} else {
		configDir = filepath.Join(dir, "configs", env)
	}

	return configDir
}

func GetConfig() *viper.Viper {
	return config
}
