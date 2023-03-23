package main

import (
	"hello-run/config"
	"hello-run/dao"
	"hello-run/router"
)

func main() {
	r := router.InitRouter()
	config.LoadEnvVariables()
	dao.InitDB()
	panic(r.Run("0.0.0.0:8080"))
}
