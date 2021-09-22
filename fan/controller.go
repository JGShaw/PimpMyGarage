package fan

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/stianeikeland/go-rpio/v4"
	"net/http"
)

type controller struct {
	relays map[string]pin
}

type pin struct {
	Active bool
	pin rpio.Pin
}

func NewController() controller {
	c := controller{
		relays: map[string]pin{
			"off": {true, rpio.Pin(5)},
			"low": {false, rpio.Pin(26)},
			"medium": {false, rpio.Pin(20)},
			"high": {false, rpio.Pin(21)},
		},
	}
	for _, pin := range c.relays {
		pin.pin.Output()
	}
	c.setAllOff()
	return c
}

func (c controller) Index(context *gin.Context) {
	context.HTML(http.StatusOK, "fan/index.tmpl", gin.H{
		"relays": c.relays,
	})
}

func (c controller) Speed(context *gin.Context) {
	c.setAllOff()
	c.setPinOn(context.Param("speed"))
	context.Redirect(303, "/fan")
}

func (c controller) setAllOff() {
	for _, pin := range c.relays {
		pin.pin.High()
		pin.active = false
	}
}

func (c controller) setPinOn(pinName string) {
	if pin, found := c.relays[pinName]; found {
		pin.pin.Low()
		pin.active = true
	}
}