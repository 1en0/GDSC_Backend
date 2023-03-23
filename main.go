package main

import (
	"hello-run/config"
	"hello-run/router"
)

func main() {
	r := router.InitRouter()
	config.LoadEnvVariables()
	config.InitDB()
	r.Run("0.0.0.0:8080")
}
