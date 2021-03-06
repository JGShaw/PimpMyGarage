package fan

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/stianeikeland/go-rpio/v4"
	"net/http"
)

type controller struct {
	relays []*pin
}

type pin struct {
	Name   string
	Active bool
	pin    rpio.Pin
}

func NewController() controller {
	c := controller{
		relays: []*pin{
			{"off", true, rpio.Pin(5)},
			{"low", false, rpio.Pin(26)},
			{"medium", false, rpio.Pin(20)},
			{"high", false, rpio.Pin(21)},
		},
	}
	for _, pin := range c.relays {
		pin.pin.Output()
	}
	c.setJustPinOn("off")
	return c
}

func (c *controller) Index(context *gin.Context) {
	context.HTML(http.StatusOK, "fan/index.tmpl", gin.H{
		"relays": c.relays,
	})
}

func (c *controller) Speed(context *gin.Context) {
	c.setJustPinOn(context.Param("speed"))
	context.Redirect(http.StatusSeeOther, "/fan")
}

func (c *controller) setJustPinOn(pinName string) {
	for _, pin := range c.relays {
		if pin.Name == pinName {
			pin.pin.Low()
			pin.Active = true
		} else {
			pin.pin.High()
			pin.Active = false

		}
	}
}
