package dao

import (
	"log"
	"strconv"
)

type Room struct {
	Id       int64  `gorm:"id"`
	RoomName string `gorm:"room_name"`
	UserId   string `gorm:"user_id"`
	City     string `gorm:"city"`
	Deleted  bool   `gorm:"deleted"`
}

func (Room) TableName() string {
	return "room"
}

func GetRoomListByUser(userId string) ([]Room, error) {
	var roomList []Room
	if err := Db.
		Model(&Room{}).
		//Where(&Room{UserId: userId, Deleted: 0}).
		Where(map[string]interface{}{"user_id": userId, "deleted": false}).
		Find(&roomList).Error; err != nil {
		log.Println("fail to get room list by user: " + userId)
		return nil, err
	}
	return roomList, nil
}

func GetRoomInfoById(id int64) (*Room, error) {
	var room Room
	if err := Db.
		Model(&Room{}).
		Where(map[string]interface{}{"id": id, "deleted": false}).
		First(&room).Error; err != nil {
		log.Println("fail to find room: " + strconv.FormatInt(id, 10))
		return nil, err
	}
	return &room, nil
}
