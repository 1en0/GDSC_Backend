package resp

import "time"

type Response[T UserVo | RecordVo | []RecordVo | FullRoomVo] struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	Comment    T      `json:"comment"`
}

type UserVo struct {
	Id     int64         `json:"id"`
	Name   string        `json:"name"`
	Avatar string        `json:"avatar"`
	Rooms  []ShortRoomVo `json:"rooms"`
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
	Age        int32 `json:"age"`
	Height     int32 `json:"height"`
	Wheelchair bool  `json:"wheelchair"`
}

type FullRoomVo struct {
	ShortRoomVo
	City      string        `json:"city"`
	Household []HouseholdVo `json:"household"`
}
