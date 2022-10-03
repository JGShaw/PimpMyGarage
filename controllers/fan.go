package controllers

import (
	"PimpMyGarage/services"
	_ "embed"
	"github.com/gin-gonic/gin"
	"net/http"
)

type fanController struct {
	fanSpeed *services.FanSpeed
}

func NewFanController(fanSpeed *services.FanSpeed) fanController {
	c := fanController{
		fanSpeed: fanSpeed,
	}
	return c
}

func (c *fanController) Index(context *gin.Context) {
	context.HTML(http.StatusOK, "fan/index.tmpl", gin.H{
		"relays": c.fanSpeed.Relays,
	})
}

func (c *fanController) Speed(context *gin.Context) {
	c.fanSpeed.SetJustPinOn(context.Param("speed"))
	context.Redirect(http.StatusSeeOther, "/fan")
}
