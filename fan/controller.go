package fan

import (
	"github.com/gin-gonic/gin"
	"github.com/stianeikeland/go-rpio/v4"
)

type controller struct {
	relays map[string]rpio.Pin
}

func NewFanController() controller {
	c := controller{
		relays: map[string]rpio.Pin{
			"1": rpio.Pin(26),
			"2": rpio.Pin(20),
			"3": rpio.Pin(21),
		},
	}
	for _, pin := range c.relays {
		pin.Output()
	}
	return c
}

func (c controller) Index(context *gin.Context) {
	context.File("PimpMyGarage/fan/index.html")
}

func (c controller) Speed(context *gin.Context) {
	c.setAllOff()
	c.setPinOn(context.Param("speed"))
	context.Redirect(303, "/fan")
}

func (c controller) setAllOff() {
	for _, pin := range c.relays {
		pin.High()
	}
}

func (c controller) setPinOn(pinName string) {
	if pin, found := c.relays[pinName]; found {
		pin.Low()
	}
}
