package main

import (
	"go-app/router"
)

func main() {
	router := router.SetupRouter()
	router.Run(":8888")
}