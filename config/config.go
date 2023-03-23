package config

import (
	"github.com/spf13/viper"
)

func LoadEnvVariables() {
	//viper.AddConfigPath(".")
	//viper.SetConfigName("")
	//viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic("fail to read env file")
		//log.Println(err)
	}
}
