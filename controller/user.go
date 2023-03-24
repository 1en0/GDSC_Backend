package controller

import (
	"github.com/gin-gonic/gin"
	"hello-run/service"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	id := c.GetString("sub")
	username := c.GetString("name")
	picture := c.GetString("picture")
	userVo, err := service.GetUserVo(id, username, picture)
	if err != nil {
		c.JSON(http.StatusOK, Response[service.UserVo]{
			StatusCode: 1,
			StatusMsg:  err.Error(),
			Comment:    service.UserVo{},
		})
	}
	c.JSON(http.StatusOK, Response[service.UserVo]{
		StatusCode: 0,
		StatusMsg:  "test",
		Comment:    *userVo,
	})
	//c.JSON(http.StatusOK, resp.Response[resp.UserVo]{
	//	StatusCode: 0,
	//	StatusMsg:  "test",
	//	Comment: resp.UserVo{
	//		Id:     "0",
	//		Name:   "test_user",
	//		Avatar: "test.pic",
	//		Rooms: []resp.ShortRoomVo{
	//			{Id: 1, RoomName: "test_room1"},
	//			{Id: 2, RoomName: "test_room2"},
	//		},
	//	},
	//})
}
