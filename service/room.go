package service

import (
	"errors"
	"hello-run/dao"
	"log"
)

func GetRoomInfoByRoomId(roomId int64) (*FullRoomVo, error) {
	//first get room info
	room, err := dao.GetRoomInfoById(roomId)
	if err != nil {
		return nil, err
	}
	//then get household info
	var householdVoList []HouseholdVo
	householdList, err := dao.GetHouseholdListByRoomId(roomId)
	if err != nil {
		return nil, err
	}
	for _, household := range householdList {
		log.Println(household.Id)
		log.Println(household.Deleted)
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

func DeleteRoomByRoomId(userId string, roomId int64) error {
	//first test if the room exists
	_, err := dao.GetRoomInfoById(roomId)
	if err != nil {
		return err
	}
	//check if it is the last room of the user
	count, err := dao.GetRoomCountByUser(userId)
	if err != nil {
		return err
	}
	if count == 1 {
		return errors.New("could not delete the last room")
	}
	return dao.DeleteRoomById(roomId)
}

type HouseholdReq struct {
	Age        int  `json:"age"`
	Height     int  `json:"height"`
	Wheelchair bool `json:"wheelchair"`
}

func CreateRoom(userId string, roomName string, city string, households []HouseholdReq) (*FullRoomVo, error) {
	roomId, err := dao.CreateRoom(userId, roomName, city)
	if err != nil {
		return nil, err
	}
	for _, household := range households {
		_, err = CreateHousehold(userId, roomId, household.Age, household.Height, household.Wheelchair)
		if err != nil {
			return nil, err
		}
	}
	//return new room info
	return GetRoomInfoByRoomId(roomId)
}

func UpdateRoomByRoomId(userId string, roomId int64, roomName string, city string, households []HouseholdReq) (*FullRoomVo, error) {
	//first test if the room exists
	householdList, err := dao.GetHouseholdListByRoomId(roomId)
	if err != nil {
		return nil, err
	}
	roomId, err = dao.UpdateRoom(roomId, roomName, city)
	if err != nil {
		return nil, err
	}

	// bulk deleted household.
	for _, household := range householdList {
		err = dao.DeleteHouseholdById(household.Id)
		if err != nil {
			return nil, err
		}
	}

	// Bulk add new household.
	for _, household := range households {
		_, err = CreateHousehold(userId, roomId, household.Age, household.Height, household.Wheelchair)
		if err != nil {
			return nil, err
		}
	}
	//return updated room info
	return GetRoomInfoByRoomId(roomId)
}
