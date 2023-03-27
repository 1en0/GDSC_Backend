package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hello-run/service"
	"net/http"
)

func ListRecordByRoomId(c *gin.Context) {
	roomId, err := checkValidInt64(c, "room_id", false)
	if err != nil {
		return
	}
	recordVoList, err := service.GetRecordListByRoomId(roomId)
	if err != nil {
		c.JSON(http.StatusOK, Response[*string]{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Fail to get record list of room %v : %v", roomId, err.Error()),
			Comment:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, Response[[]service.RecordVo]{
		StatusCode: 0,
		StatusMsg:  "",
		Comment:    recordVoList,
	})
}

func UpdateRecord(c *gin.Context) {
	recordId, err := checkValidInt64(c, "record_id", false)
	if err != nil {
		return
	}
	riskTypeStr := c.Query("risk_type")
	title := c.Query("title")
	content := c.Query("content")
	screenshot := c.Query("screenshot")
	recordVo, err := service.UpdateRecord(recordId, riskTypeStr, title, content, screenshot)
	if err != nil {
		c.JSON(http.StatusOK, Response[*string]{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Fail to update record %v: %v", recordId, err.Error()),
			Comment:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, Response[service.RecordVo]{
		StatusCode: 0,
		StatusMsg:  "",
		Comment:    *recordVo,
	})
}

func CreateRecord(c *gin.Context) {
	userId := c.GetString("sub")
	roomId, err := checkValidInt64(c, "room_id", false)
	if err != nil {
		return
	}
	riskTypeStr := c.Query("risk_type")
	title := c.Query("title")
	content := c.Query("content")
	screenshot := c.Query("screenshot")
	recordVo, err := service.CreateRecord(userId, roomId, riskTypeStr, title, content, screenshot)
	if err != nil {
		c.JSON(http.StatusOK, Response[*string]{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Fail to create record: %v", err.Error()),
			Comment:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, Response[service.RecordVo]{
		StatusCode: 0,
		StatusMsg:  "",
		Comment:    *recordVo,
	})
}

func DeleteRecord(c *gin.Context) {
	recordId, err := checkValidInt64(c, "record_id", false)
	if err != nil {
		return
	}
	err = service.DeleteRecord(recordId)
	if err != nil {
		c.JSON(http.StatusOK, Response[*string]{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Fail to delete record: %v :%v", recordId, err.Error()),
			Comment:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, Response[*string]{
		StatusCode: 0,
		StatusMsg:  fmt.Sprintf("Successfully deleted record: %v", recordId),
		Comment:    nil,
	})
}
