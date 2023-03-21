package main

import (
	"hello-run/router"
)

func main() {
	r := router.InitRouter()

	r.Run("0.0.0.0:8080")
}
