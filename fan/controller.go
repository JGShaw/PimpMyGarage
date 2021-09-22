package fan

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/stianeikeland/go-rpio/v4"
	"net/http"
)

type controller struct {
	relays map[string]rpio.Pin
}

func NewController() controller {
	c := controller{
		relays: map[string]rpio.Pin{
			"off": rpio.Pin(5),
			"low": rpio.Pin(26),
			"medium": rpio.Pin(20),
			"high": rpio.Pin(21),
		},
	}
	for _, pin := range c.relays {
		pin.Output()
	}
	c.setAllOff()
	return c
}

func (c controller) Index(context *gin.Context) {
	context.HTML(http.StatusOK, "fan/index.tmpl", gin.H{
		"relays": c.relays,
		"active": c.currentActive(),
	})
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

func (c controller) currentActive() string {
	for key, pin := range c.relays {
		if pin.Read() == rpio.Low {
			return key
		}
	}
	return "0"
}
