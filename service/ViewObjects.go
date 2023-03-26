package service

import "time"

type UserVo struct {
	Id      string        `json:"id"`
	Name    string        `json:"name"`
	Picture string        `json:"picture"`
	Rooms   []ShortRoomVo `json:"rooms"`
}

type ShortRoomVo struct {
	Id       int64  `json:"id"`
	RoomName string `json:"room_name"`
}

type RecordVo struct {
	Id         int64     `json:"id"`
	RiskType   string    `json:"risk_type"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Screenshot string    `json:"screenshot"`
	CreateTime time.Time `json:"create_time"`
}

type HouseholdVo struct {
	Id         int64 `json:"id"`
	Age        int   `json:"age"`
	Height     int   `json:"height"`
	Wheelchair bool  `json:"wheelchair"`
}

type FullRoomVo struct {
	ShortRoomVo
	City      string        `json:"city"`
	Household []HouseholdVo `json:"household"`
}
