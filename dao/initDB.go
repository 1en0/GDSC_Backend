package dao

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() {
	path := os.Getenv("DB_PATH")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	arg := os.Getenv("DB_ARG")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", username, password, path, port, dbName, arg)

	fmt.Println(args)

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("fail to connect database")
	}
	Db = db
}
