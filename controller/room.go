package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hello-run/service"
	"net/http"
	"strconv"
)

func GetRoomInfo(c *gin.Context) {
	//id := c.GetString("sub")
	roomId, err := strconv.ParseInt(c.Query("room_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response[service.FullRoomVo]{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Room id %v is invalid: %v", roomId, err.Error()),
		})
		return
	}

	roomVo, err := service.GetRoomInfoByRoomId(roomId)
	if err != nil {
		c.JSON(http.StatusOK, Response[service.FullRoomVo]{
			StatusCode: 1,
			//StatusMsg:  "fail to get info of room:" + c.Query("room_id") + err.Error(),
			StatusMsg: fmt.Sprintf("Fail to get info of room: %v : %v ", roomId, err.Error()),
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
	roomId, err := strconv.ParseInt(c.Query("room_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response[service.FullRoomVo]{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Room id %v is invalid: %v", roomId, err.Error()),
		})
		return
	}
	roomName := c.Query("room_name")
	city := c.Query("city")
	roomVo, err := service.UpdateRoomByRoomId(roomId, roomName, city)
	if err != nil {
		c.JSON(http.StatusOK, Response[service.FullRoomVo]{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Fail to update room: %v :%v", roomId, err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, Response[service.FullRoomVo]{
		StatusCode: 0,
		StatusMsg:  "",
		Comment:    *roomVo,
	})
}

func CreateRoom(c *gin.Context) {
	id := c.GetString("sub")
	roomName := c.Query("room_name")
	city := c.Query("city")
	roomVo, err := service.CreateRoom(id, roomName, city)
	if err != nil {
		c.JSON(http.StatusOK, Response[service.FullRoomVo]{
			StatusCode: 1,
			StatusMsg:  "Fail to create room: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response[service.FullRoomVo]{
		StatusCode: 0,
		StatusMsg:  "",
		Comment:    *roomVo,
	})
}

func DeleteRoom(c *gin.Context) {
	roomId, _ := strconv.ParseInt(c.Query("room_id"), 10, 64)
	err := service.DeleteRoomByRoomId(roomId)
	if err != nil {
		c.JSON(http.StatusOK, Response[string]{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Fail to delete room: %v :%v", roomId, err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, Response[string]{
		StatusCode: 0,
		StatusMsg:  fmt.Sprintf("Successfully deleted room: %v", roomId),
	})
}
