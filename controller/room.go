package controller

import (
	"github.com/gin-gonic/gin"
	"hello-run/service"
	"net/http"
	"strconv"
)

func GetRoomInfo(c *gin.Context) {
	//id := c.GetString("sub")
	roomId, _ := strconv.ParseInt(c.Query("room_id"), 10, 64)
	roomVo, err := service.GetRoomInfoByRoomId(roomId)
	if err != nil {
		c.JSON(http.StatusOK, Response[service.FullRoomVo]{
			StatusCode: 1,
			StatusMsg:  err.Error(),
			Comment:    service.FullRoomVo{},
		})
		return
	}
	c.JSON(http.StatusOK, Response[service.FullRoomVo]{
		StatusCode: 0,
		StatusMsg:  "",
		Comment:    *roomVo,
	})
}

func UpdateRoom(c *gin.Context) {

}

func CreateRoom(c *gin.Context) {

}

func DeleteRoom(c *gin.Context) {

}
