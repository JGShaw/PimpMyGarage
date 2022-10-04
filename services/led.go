package services

import (
	"github.com/hisamafahri/coco"
	magichome "github.com/moonliightz/magic-home/pkg"
	"net"
)

type ledService struct {
	controller *magichome.Controller
	minHue     float64
	maxHue     float64
}

func NewLedService() *ledService {
	c, _ := magichome.New(net.ParseIP("192.168.1.104"), 5577)
	return &ledService{
		controller: c,
		minHue:     240,
		maxHue:     0,
	}
}

func (l *ledService) SetColorPercentage(percentage float64) error {
	hue := l.minHue + ((l.maxHue - l.minHue) * percentage)
	rgb := coco.Hsl2Rgb(hue, 100, 50)
	return l.controller.SetColor(magichome.Color{
		R: uint8(rgb[0]),
		G: uint8(rgb[1]),
		B: uint8(rgb[2]),
		W: 0,
	})
}
