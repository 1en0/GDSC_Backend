package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hello-run/dao/cloudsql"
)

var Db *gorm.DB

func InitDB() {
	//path := os.Getenv("DB_PATH")
	//dbName := os.Getenv("DB_NAME")
	//port := os.Getenv("DB_PORT")
	//username := os.Getenv("DB_USERNAME")
	//password := os.Getenv("DB_PASSWORD")
	//arg := os.Getenv("DB_ARG")
	//args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", username, password, path, port, dbName, arg)
	//
	//fmt.Println(args)

	//db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
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

	//db.AutoMigrate(
	//	&User{},
	//	&Record{},
	//	&Room{},
	//	&Household{},
	//)

	Db = db
}
