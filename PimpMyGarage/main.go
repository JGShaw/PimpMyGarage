package main

import (
	"PimpMyGarage/PimpMyGarage/fan"
	"github.com/gin-gonic/gin"
	"github.com/stianeikeland/go-rpio/v4"

)

var relay1 rpio.Pin
var relay2 rpio.Pin
var relay3 rpio.Pin

func main() {
	err := rpio.Open()
	if err != nil {
		panic(err)
	}
	defer  rpio.Close()

	fanController := fan.NewFanController()

	router := gin.Default()
	router.GET("/fan", fanController.Index)
	router.GET("/fan/:speed", fanController.Speed)

	router.Run(":8080")
}