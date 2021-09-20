package main

import (
	"PimpMyGarage/fan"
	"github.com/gin-gonic/gin"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	err := rpio.Open()
	if err != nil {
		panic(err)
	}
	defer rpio.Close()

	fanController := fan.NewFanController()

	router := gin.Default()
	router.GET("/fan", fanController.Index)
	router.GET("/fan/speed/:speed", fanController.Speed)

	router.Run(":8080")
}
