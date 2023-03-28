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

func GetRoomCountByUser(userId string) (int64, error) {
	var count int64
	if err := Db.
		Model(&Room{}).
		Where(map[string]interface{}{"user_id": userId, "deleted": false}).
		Count(&count).
		Error; err != nil {
		return 0, err
	}
	return count, nil
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

// if no room satisfies the search condition
// gorm.ErrRecordNotFound will be thrown

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

func DeleteRoomById(id int64) error {
	if err := Db.
		Model(&Room{}).
		Where(map[string]interface{}{"id": id}).
		Update("deleted", true).Error; err != nil {
		log.Println("fail to delete room: " + strconv.FormatInt(id, 10))
		return err
	}
	return nil
}

func CreateRoom(userId string, roomName string, city string) (int64, error) {
	room := Room{
		RoomName: roomName,
		UserId:   userId,
		City:     city,
		Deleted:  false,
	}
	if err := Db.
		Model(&Room{}).
		Create(&room).Error; err != nil {
		return 0, err
	}
	return room.Id, nil
}

func UpdateRoom(id int64, roomName string, city string) (int64, error) {
	if err := Db.
		Model(&Room{Id: id}).
		Updates(Room{RoomName: roomName, City: city}).
		Error; err != nil {
		return 0, err
	}
	return id, nil
}
