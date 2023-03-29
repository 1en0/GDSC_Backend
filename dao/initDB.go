package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hello-run/dao/cloudsql"
)

var Db *gorm.DB

func InitDB() {
	//path := viper.GetString("DB_PATH")
	//dbName := viper.GetString("DB_NAME")
	//port := viper.GetInt64("DB_PORT")
	//username := viper.GetString("DB_USERNAME")
	//password := viper.GetString("DB_PASSWORD")
	//arg := viper.GetString("DB_ARG")
	//args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", username, password, path, port, dbName, arg)
	cloudSQL := cloudsql.GetDB()
	db, err := gorm.Open(mysql.New(mysql.Config{Conn: cloudSQL}), &gorm.Config{})
	if err != nil {
		panic("fail to connect database")
	}
	Db = db
}
