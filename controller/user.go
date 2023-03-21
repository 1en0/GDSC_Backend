package controller

import (
	"github.com/gin-gonic/gin"
	"hello-run/controller/resp"
	"net/http"
)

func UserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, resp.Response[resp.UserVo]{
		StatusCode: 0,
		StatusMsg:  "test",
		Comment: resp.UserVo{
			Id:     0,
			Name:   "test_user",
			Avatar: "test.pic",
			Rooms: []resp.ShortRoomVo{
				{Id: 1, RoomName: "test_room1"},
				{Id: 2, RoomName: "test_room2"},
			},
		},
	})
}
