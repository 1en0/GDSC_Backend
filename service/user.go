package service

import "hello-run/dao"

func isNewUser(userid string) (bool, error) {
	exists, err := dao.IfUserExists(userid)
	if err != nil {
		return false, err
	}
	if exists {
		return false, nil
	} else {
		err = dao.CreateUser(userid)
		if err != nil {
			return false, err
		}
		return true, nil
	}
}

func GetShortRoomInfoList(userid string) ([]ShortRoomVo, error) {
	roomList, err := dao.GetRoomListByUser(userid)
	if err != nil {
		return nil, err
	}
	var infoList []ShortRoomVo
	for _, room := range roomList {
		infoList = append(infoList, ShortRoomVo{Id: room.Id, RoomName: room.RoomName})
	}
	return infoList, nil
}

func GetUserVo(userid string, username string, picture string) (*UserVo, error) {
	isNew, err := isNewUser(userid)
	if err != nil {
		return nil, err
	}
	if isNew {
		return &UserVo{IsNewUser: true}, nil
	} else {
		infoList, err := GetShortRoomInfoList(userid)
		if err != nil {
			return nil, err
		}
		userVo := UserVo{
			Id:        userid,
			Name:      username,
			Picture:   picture,
			Rooms:     infoList,
			IsNewUser: false,
		}
		return &userVo, nil
	}
}
