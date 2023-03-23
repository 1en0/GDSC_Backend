package service

import "hello-run/dao"

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
	infoList, err := GetShortRoomInfoList(userid)
	if err != nil {
		return nil, err
	}
	userVo := UserVo{
		Id:      userid,
		Name:    username,
		Picture: picture,
		Rooms:   infoList,
	}
	return &userVo, nil
}
