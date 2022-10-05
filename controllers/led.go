package controllers

import (
	"PimpMyGarage/services"
	_ "embed"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

type ledController struct {
	ipAddress  string
	port       int
	ledService *services.LedService
	hrmService *services.HrmService
}

func NewLedController(hrmService *services.HrmService) ledController {
	c := ledController{
		ipAddress:  "192.168.1.104",
		port:       5577,
		hrmService: hrmService,
	}
	c.setLedService()
	return c
}

func (l *ledController) Index(context *gin.Context) {
	//TODO sanatize and check for errors
	if ip, found := context.GetQuery("ipAddress"); found {
		l.ipAddress = ip
		l.setLedService()
	}
	if port, found := context.GetQuery("port"); found {
		l.port, _ = strconv.Atoi(port)
		l.setLedService()
	}

	context.HTML(http.StatusOK, "led/index.tmpl", gin.H{
		"ipAddress": l.ipAddress,
		"port":      l.port,
	})
}

func (l *ledController) setLedService() {
	if l.ledService != nil {
		l.ledService.Close()
	}
	l.ledService = services.NewLedService(l.ipAddress, uint16(l.port))
	l.hrmService.AddListeners(
		[]func(float64){
			func(hr float64) {
				hrMin := 110.0
				hrMax := 185.0
				hr = math.Max(math.Min(hr, hrMax), hrMin)
				percentage := (hr - hrMin) / (hrMax - hrMin)
				err := l.ledService.SetColorPercentage(percentage)
				if err != nil {
					return
				}
			},
		},
	)

}
