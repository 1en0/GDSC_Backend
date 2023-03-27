package dao

import (
	"fmt"
	"hello-run/constant"
	"log"
	"time"
)

type Record struct {
	Id         int64             `gorm:"id"`
	UserId     string            `gorm:"user_id"`
	RoomId     int64             `gorm:"room_id"`
	RiskType   constant.RiskType `gorm:"risk_type"`
	Title      string            `gorm:"title"`
	Content    string            `gorm:"content"`
	Screenshot string            `gorm:"screenshot"`
	Deleted    bool              `gorm:"deleted"`
	CreatedAt  time.Time         `json:"created_at"`
}

func (Record) TableName() string {
	return "record"
}

func CreateRecord(userId string, roomId int64, riskType constant.RiskType, title string, content string, screenshot string) (*Record, error) {
	record := Record{
		UserId:     userId,
		RoomId:     roomId,
		RiskType:   riskType,
		Title:      title,
		Content:    content,
		Screenshot: screenshot,
		Deleted:    false,
	}
	if err := Db.
		Model(&Record{}).
		Create(&record).
		Error; err != nil {
		return nil, err
	}
	return &record, nil
}

func DeleteRecordById(id int64) error {
	if err := Db.
		Model(&Record{}).
		Where(map[string]interface{}{"id": id}).
		Update("deleted", true).
		Error; err != nil {
		log.Println(fmt.Sprintf("fail to delete room: %v", id))
		return err
	}
	return nil
}

func UpdateRecord(id int64, riskType constant.RiskType, title string, content string, screenshot string) (int64, error) {
	if err := Db.
		Model(&Record{Id: id}).
		Updates(Record{RiskType: riskType, Title: title, Content: content, Screenshot: screenshot}).
		Error; err != nil {
		return 0, err
	}
	return id, nil
}

func GetRecordListByRoomId(roomId int64) ([]Record, error) {
	var recordList []Record
	if err := Db.
		Model(&Record{}).
		Where(map[string]interface{}{"room_id": roomId, "deleted": false}).
		Find(&recordList).
		Error; err != nil {
		log.Println(fmt.Sprintf("Fail to get record list by room id: %v", roomId))
		return nil, err
	}
	return recordList, nil
}

func GetRecordById(id int64) (*Record, error) {
	var record Record
	if err := Db.
		Model(&Record{}).
		Where(map[string]interface{}{"id": id, "deleted": false}).
		First(&record).Error; err != nil {
		log.Println(fmt.Sprintf("fail to find record: %v", id))
		return nil, err
	}
	return &record, nil
}
