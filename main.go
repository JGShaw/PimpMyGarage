package main

import (
	"PimpMyGarage/fan"
	"github.com/gin-gonic/gin"
)

func main() {
	//err := rpio.Open()
	//if err != nil {
	//	panic(err)
	//}
	//defer rpio.Close()

	fanController := fan.NewController()

	router := gin.Default()
	router.GET("/fan", fanController.Index)
	router.GET("/fan/speed/:speed", fanController.Speed)

	router.Run(":8080")
}