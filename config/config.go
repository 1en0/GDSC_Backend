package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnvVariables() {
	//viper.AddConfigPath(".")
	//viper.SetConfigName("")
	//viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading env file")
		//log.Println(err)
	}
}

func InitDB() {
	path := viper.GetString("DB_PATH")
	dbName := viper.GetString("DB_NAME")
	port := viper.GetInt64("DB_PORT")
	username := viper.GetString("DB_USERNAME")
	password := viper.GetString("DB_PASSWORD")
	arg := viper.GetString("DB_ARG")
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", username, password, path, port, dbName, arg)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("fail to connect database")
	}
	DB = db
}
