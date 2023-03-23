package dao

import (
	"log"
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
		log.Println("fail to get room list by user")
		return nil, err
	}
	return roomList, nil
}
