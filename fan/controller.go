package fan

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/stianeikeland/go-rpio/v4"
)

//go:embed index.html
var indexPage string

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
	context.Data(200, "text/plain; charset=utf-8", []byte(indexPage))
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
