package router

import (
	"github.com/gin-gonic/gin"
	"hello-run/controller"
	"hello-run/middleware/googleauth"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(googleauth.FakeAuth())

	router.GET("/user/", controller.UserInfo)

	router.GET("/record/list/", controller.ListRecord)
	router.POST("/record/create/", controller.CreateRecord)
	router.POST("/record/update/", controller.UpdateRecord)
	router.POST("/record/delete/", controller.DeleteRecord)

	router.GET("/room/", controller.RoomInfo)
	router.POST("/room/update/", controller.UpdateRoom)
	router.POST("/room/create/", controller.CreateRoom)
	router.POST("/room/delete/", controller.DeleteRoom)

	router.POST("/household/create/", controller.CreateHousehold)
	router.POST("/household/update/", controller.UpdateHousehold)
	router.POST("/household/delete/", controller.DeleteHousehold)

	return router
}
