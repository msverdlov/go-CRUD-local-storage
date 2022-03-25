package config

import (
	"github.com/spf13/viper"
	"log"
)

const (
	ConfFilePath = "../../config"
	ConfFileName = "config"
	ConfFileExt  = "yaml"
)

func SetConfig(filePath string, fileName string, configType string) {
	viper.AddConfigPath(filePath)
	viper.SetConfigName(fileName)
	viper.SetConfigType(configType)
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}
}

func GetViperValueByKey(key string) string {
	return viper.GetString(key)
}
