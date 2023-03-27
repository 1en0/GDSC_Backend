package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func checkValidInt64(c *gin.Context, paramKey string, optional bool) (int64, error) {
	value := c.Query(paramKey)
	if optional && value == "" {
		return 0, nil
	}
	id, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response[*string]{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("%v : %v is invalid: %v", paramKey, value, err.Error()),
			Comment:    nil,
		})
		return 0, err
	}
	return id, nil
}

func checkValidInt(c *gin.Context, paramKey string, optional bool) (int, error) {
	value := c.Query(paramKey)
	if optional && value == "" {
		return 0, nil
	}
	param, err := strconv.Atoi(value)
	if err != nil {
		c.JSON(http.StatusOK, Response[*string]{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("%v : %v is invalid: %v", paramKey, value, err.Error()),
			Comment:    nil,
		})
		return 0, err
	}
	return param, nil
}

func checkValidBool(c *gin.Context, paramKey string, optional bool) (bool, error) {
	value := c.Query(paramKey)
	if optional && value == "" {
		return false, nil
	}
	param, err := strconv.ParseBool(value)
	if err != nil {
		c.JSON(http.StatusOK, Response[*string]{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("%v : %v is invalid: %v", paramKey, value, err.Error()),
			Comment:    nil,
		})
		return false, err
	}
	return param, nil
}
