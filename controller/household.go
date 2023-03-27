package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hello-run/service"
	"net/http"
)

func UpdateHousehold(c *gin.Context) {
	householdId, err := checkValidInt64(c, "household_id", false)
	if err != nil {
		return
	}
	age, err := checkValidInt(c, "age", true)
	if err != nil {
		return
	}
	height, err := checkValidInt(c, "height", true)
	if err != nil {
		return
	}
	wheelchair, err := checkValidBool(c, "wheelchair", true)
	if err != nil {
		return
	}

	householdVo, err := service.UpdateHousehold(householdId, age, height, wheelchair)
	if err != nil {
		c.JSON(http.StatusOK, Response[*string]{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Fail to update room: %v :%v", householdId, err.Error()),
			Comment:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, Response[service.HouseholdVo]{
		StatusCode: 0,
		StatusMsg:  "",
		Comment:    *householdVo,
	})
}

func CreateHousehold(c *gin.Context) {
	userId := c.GetString("sub")
	roomId, err := checkValidInt64(c, "room_id", false)
	if err != nil {
		return
	}
	age, err := checkValidInt(c, "age", false)
	if err != nil {
		return
	}
	height, err := checkValidInt(c, "height", false)
	if err != nil {
		return
	}
	wheelchair, err := checkValidBool(c, "wheelchair", false)
	if err != nil {
		return
	}

	householdVo, err := service.CreateHousehold(userId, roomId, age, height, wheelchair)
	if err != nil {
		c.JSON(http.StatusOK, Response[*string]{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Fail to create household: %v", err.Error()),
			Comment:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, Response[service.HouseholdVo]{
		StatusCode: 0,
		StatusMsg:  "",
		Comment:    *householdVo,
	})
}

func DeleteHousehold(c *gin.Context) {
	householdId, err := checkValidInt64(c, "household_id", false)
	if err != nil {
		return
	}
	err = service.DeleteHouseholdByHouseholdId(householdId)
	if err != nil {
		c.JSON(http.StatusOK, Response[*string]{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("Fail to delete household: %v :%v", householdId, err.Error()),
			Comment:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, Response[*string]{
		StatusCode: 0,
		StatusMsg:  fmt.Sprintf("Successfully deleted household: %v", householdId),
		Comment:    nil,
	})

}
