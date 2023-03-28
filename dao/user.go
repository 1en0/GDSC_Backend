package dao

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	Id string `gorm:"id"`
}

func (User) TableName() string {
	return "user"
}

func CreateUser(id string) error {
	if err := Db.Create(&User{Id: id}).Error; err != nil {
		log.Println(fmt.Sprintf("Fail to create user: %v", id))
		return err
	}
	return nil
}

func IfUserExists(id string) (bool, error) {
	err := Db.First(&User{Id: id}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	} else {
		return true, nil
	}
}
