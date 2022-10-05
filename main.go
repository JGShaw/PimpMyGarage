package main

import (
	"PimpMyGarage/controllers"
	"PimpMyGarage/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stianeikeland/go-rpio/v4"
	"math"
)

func main() {
	err := rpio.Open()
	if err != nil {
		panic(err)
	}
	defer rpio.Close()

	fanSpeedService := services.NewFanSpeed()
	hrmService, _ := services.NewHrmService()
	ledService := services.NewLedService()

	rootController := controllers.NewRootController()
	fanController := controllers.NewFanController(fanSpeedService)
	hrmController := controllers.NewHrmController(hrmService, []func(float64){
		func(hr float64) {
			fmt.Println(hr)
			hrMin := 60.0
			hrMax := 150.0
			hr = math.Max(math.Min(hr, hrMax), hrMin)
			percentage := (hr - hrMin) / (hrMax - hrMin)
			err := ledService.SetColorPercentage(percentage)
			if err != nil {
				return
			}
		},
	})
	ledController := controllers.NewLedController()

	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.StaticFile("templates/pimp_my_garage.svg", "./templates/pimp_my_garage.svg")

	router.GET("/", rootController.Index)

	router.GET("/fan", fanController.Index)
	router.GET("/fan/speed/:speed", fanController.Speed)

	router.GET("/hrm", hrmController.Index)
	router.GET("/hrm/search", hrmController.Search)
	router.GET("/hrm/connect/:address", hrmController.Connect)
	router.GET("/hrm/disconnect", hrmController.Disconnect)

	router.GET("/led", ledController.Index)

	defer hrmService.Disconnect()
	router.Run(":8080")
}
