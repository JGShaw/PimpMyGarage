package services

import "github.com/stianeikeland/go-rpio/v4"

type pin struct {
	Name   string
	Active bool
	pin    rpio.Pin
}

type FanSpeed struct {
	Relays []*pin
}

func NewFanSpeed() *FanSpeed {
	f := &FanSpeed{
		Relays: []*pin{
			{"off", true, rpio.Pin(5)},
			{"low", false, rpio.Pin(26)},
			{"medium", false, rpio.Pin(20)},
			{"high", false, rpio.Pin(21)},
		},
	}
	for _, pin := range f.Relays {
		pin.pin.Output()
	}
	f.SetJustPinOn("off")
	return f
}

func (f *FanSpeed) SetJustPinOn(pinName string) {
	for _, pin := range f.Relays {
		if pin.Name == pinName {
			pin.pin.Low()
			pin.Active = true
		} else {
			pin.pin.High()
			pin.Active = false

		}
	}
}
