package dao

import (
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
