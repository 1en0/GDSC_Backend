package dao

import (
	"fmt"
	"log"
	"strconv"
)

type Household struct {
	Id         int64  `gorm:"id"`
	UserId     string `gorm:"user_id"`
	RoomId     int64  `gorm:"room_id"`
	Age        int    `gorm:"age"`
	Height     int    `gorm:"height"`
	Wheelchair bool   `gorm:"wheelchair"`
	Deleted    bool   `gorm:"deleted"`
}

func (Household) TableName() string {
	return "household"
}

//if no household satisfies the search condition
//it will return a list of Household with length 0
//no error will be thrown

func GetHouseholdListByRoomId(roomId int64) ([]Household, error) {
	var householdList []Household
	if err := Db.
		Model(&Household{}).
		Where(map[string]interface{}{"room_id": roomId, "deleted": false}).
		Find(&householdList).Error; err != nil {
		log.Println("fail to get household list by room id:" + strconv.FormatInt(roomId, 10))
		return nil, err
	}
	return householdList, nil
}

func CreateHousehold(userId string, roomId int64, age int, height int, wheelchair bool) (*Household, error) {
	household := Household{
		UserId:     userId,
		RoomId:     roomId,
		Age:        age,
		Height:     height,
		Wheelchair: wheelchair,
		Deleted:    false,
	}
	if err := Db.
		Model(&Household{}).
		Create(&household).
		Error; err != nil {
		return nil, err
	}
	return &household, nil
}

func UpdateHousehold(id int64, age int, height int, wheelchair bool) (int64, error) {
	if err := Db.
		Model(&Household{Id: id}).
		Updates(Household{Age: age, Height: height, Wheelchair: wheelchair}).
		Error; err != nil {
		return 0, err
	}
	return id, nil
}

func DeleteHouseholdById(id int64) error {
	if err := Db.
		Model(&Household{}).
		Where(map[string]interface{}{"id": id}).
		Update("deleted", true).
		Error; err != nil {
		log.Println(fmt.Sprintf("fail to delete room: %v", id))
		return err
	}
	return nil
}

// if no room satisfies the search condition
// gorm.ErrRecordNotFound will be thrown

func GetHouseholdInfoById(id int64) (*Household, error) {
	var household Household
	if err := Db.
		Model(&Household{}).
		Where(map[string]interface{}{"id": id, "deleted": false}).
		First(&household).Error; err != nil {
		log.Println(fmt.Sprintf("fail to find room: %v", id))
		return nil, err
	}
	return &household, nil
}
