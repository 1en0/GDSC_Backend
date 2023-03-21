package config

//
//import (
//	"fmt"
//	"github.com/jinzhu/gorm"
//	"github.com/spf13/viper"
//)
//
//var Db *gorm.DB
//
//func InitDb() {
//	driverName := viper.GetString("db.driver_name")
//	host := viper.GetString("db.host")
//	port := viper.GetString("db.port")
//	database := viper.GetString("db.database")
//	username := viper.GetString("db.username")
//	password := viper.GetString("db.password")
//	charset := viper.GetString("db.charset")
//	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
//		username,
//		password,
//		host,
//		port,
//		database,
//		charset,
//	)
//
//	//db, err := gorm.Open(driverName, args)
//
//	if err != nil {
//		panic("failed to connect database,err:" + err.Error())
//	}
//
//	Db = db
//}
