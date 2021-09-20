package main

import (
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

	setupRelays()

	router := gin.Default()
	router.GET("/fan/:speed", fanSpeed)

	router.Run(":8080")
}

func setupRelays() {
	relay1 = rpio.Pin(26)
	relay2 = rpio.Pin(20)
	relay3 = rpio.Pin(21)
	relay1.Output()
	relay2.Output()
	relay3.Output()
}

func fanSpeed(c *gin.Context) {
	relay1.Low()
	relay2.Low()
	relay3.Low()

	switch c.Param("speed") {
	case "1":
		relay1.High()
	case "2":
		relay2.High()
	case "3":
		relay3.High()
	}
}
