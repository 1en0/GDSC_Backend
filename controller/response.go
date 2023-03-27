package controller

import "hello-run/service"

type Response[T service.UserVo | service.RecordVo | []service.RecordVo | service.FullRoomVo | string | service.HouseholdVo | *string] struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	Comment    T      `json:"comment"`
}
