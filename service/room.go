package service

import (
	"errors"
	"gorm.io/gorm"
	"hello-run/dao"
)

func GetRoomInfoByRoomId(roomId int64) (*FullRoomVo, error) {
	room, err := dao.GetRoomInfoById(roomId)
	if err != nil {
		return nil, err
	}

	var householdVoList []HouseholdVo
	householdList, err := dao.GetHouseholdListByRoomId(roomId)

	//if household in this room cannot be found
	//householdVoList is an empty list
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			householdVoList = []HouseholdVo{}
		} else {
			return nil, err
		}
	}

	householdVoList = GetHouseholdVoList(householdList)

	shortRoomVo := ShortRoomVo{
		Id:       room.Id,
		RoomName: room.RoomName,
	}

	return &FullRoomVo{
		ShortRoomVo: shortRoomVo,
		City:        room.City,
		Household:   householdVoList,
	}, nil
}
