package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hello-run/service"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	id := c.GetString("sub")
	userVo, err := service.GetUserVo(id)
	if err != nil {
		c.JSON(http.StatusOK, Response[*string]{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Fail to get information of user %v: %v", id, err.Error()),
			Comment:    nil,
		})
	}
	c.JSON(http.StatusOK, Response[service.UserVo]{
		StatusCode: 0,
		StatusMsg:  "",
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
