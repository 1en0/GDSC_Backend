package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"hello-run/service"
	"net/http"
)

func GetRoomInfo(c *gin.Context) {
	//id := c.GetString("sub")
	roomId, err := checkValidInt64(c, "room_id", false)
	if err != nil {
		return
	}

	roomVo, err := service.GetRoomInfoByRoomId(roomId)
	if err != nil {
		c.JSON(http.StatusOK, Response[*string]{
			StatusCode: 1,
			//StatusMsg:  "fail to get info of room:" + c.Query("room_id") + err.Error(),
			StatusMsg: fmt.Sprintf("Fail to get info of room: %v : %v ", roomId, err.Error()),
			Comment:   nil,
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
	userId := c.GetString("sub")
	roomId, err := checkValidInt64(c, "room_id", false)
	if err != nil {
		return
	}

	roomName := c.Query("room_name")
	city := c.Query("city")
	data, err := c.GetRawData()
	var households []service.HouseholdReq
	err = json.Unmarshal(data, &households)
	if err != nil {
		c.JSON(http.StatusOK, Response[*string]{
			StatusCode: 1,
			StatusMsg:  "Wrong format for household " + err.Error(),
			Comment:    nil,
		})
		return
	}
	roomVo, err := service.UpdateRoomByRoomId(userId, roomId, roomName, city, households)
	if err != nil {
		c.JSON(http.StatusOK, Response[*string]{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Fail to update room: %v :%v", roomId, err.Error()),
			Comment:    nil,
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
	userId := c.GetString("sub")
	roomName := c.Query("room_name")
	city := c.Query("city")
	data, err := c.GetRawData()
	var households []service.HouseholdReq
	err = json.Unmarshal(data, &households)
	if err != nil {
		c.JSON(http.StatusOK, Response[*string]{
			StatusCode: 1,
			StatusMsg:  "Wrong format for household " + err.Error(),
			Comment:    nil,
		})
		return
	}
	roomVo, err := service.CreateRoom(userId, roomName, city, households)
	if err != nil {
		c.JSON(http.StatusOK, Response[*string]{
			StatusCode: 1,
			StatusMsg:  "Fail to create room: " + err.Error(),
			Comment:    nil,
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
	userId := c.GetString("sub")
	roomId, err := checkValidInt64(c, "room_id", false)
	if err != nil {
		return
	}
	err = service.DeleteRoomByRoomId(userId, roomId)
	if err != nil {
		c.JSON(http.StatusOK, Response[*string]{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Fail to delete room: %v :%v", roomId, err.Error()),
			Comment:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, Response[*string]{
		StatusCode: 0,
		StatusMsg:  fmt.Sprintf("Successfully deleted room: %v", roomId),
		Comment:    nil,
	})
}
