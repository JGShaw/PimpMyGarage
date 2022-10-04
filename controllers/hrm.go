package controllers

import (
	"PimpMyGarage/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type hrmController struct {
	hrmService *services.HrmService
	listeners  []func(float64)
}

func NewHrmController(service *services.HrmService, listeners []func(float64)) hrmController {
	return hrmController{
		hrmService: service,
		listeners:  listeners,
	}
}

func (c *hrmController) Index(context *gin.Context) {
	context.HTML(http.StatusOK, "hrm/index.tmpl", gin.H{
		"hrm": c.hrmService.ConnectedDeviceName(),
	})
}

func (c *hrmController) Search(context *gin.Context) {
	found, err := c.hrmService.Scan()
	if err != nil {
		fmt.Println(err)
	}
	context.HTML(http.StatusOK, "hrm/search.tmpl", gin.H{
		"found": found,
	})
}

func (c *hrmController) Connect(context *gin.Context) {
	err := c.hrmService.Connect(context.Param("address"))
	if err != nil {
		fmt.Println(err)
	}
	c.hrmService.AddListeners(c.listeners)
	context.Redirect(http.StatusSeeOther, "/hrm")
}

func (c *hrmController) Disconnect(context *gin.Context) {
	err := c.hrmService.Disconnect()
	if err != nil {
		fmt.Println(err)
	}
	context.Redirect(http.StatusSeeOther, "/hrm")
}
