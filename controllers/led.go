package controllers

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ledController struct {
	ipAddress string
	port      int
}

func NewLedController() ledController {
	c := ledController{
		ipAddress: "192.168.1.104",
		port:      5577,
	}
	return c
}

func (l *ledController) Index(context *gin.Context) {
	//TODO sanatize and check for errors
	if ip, found := context.GetQuery("ipAddress"); found {
		l.ipAddress = ip
	}
	if port, found := context.GetQuery("port"); found {
		l.port, _ = strconv.Atoi(port)
	}

	context.HTML(http.StatusOK, "led/index.tmpl", gin.H{
		"ipAddress": l.ipAddress,
		"port":      l.port,
	})
}
