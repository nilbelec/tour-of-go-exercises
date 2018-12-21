package main

import (
	"image"
	"image/color"
	"math"

	"golang.org/x/tour/pic"
)

type myImage struct{}

func (i myImage) Bounds() image.Rectangle {
	h, w := 100, 100
	return image.Rect(0, 0, w, h)
}

func (i myImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (i myImage) At(x, y int) color.Color {
	v := uint8(y - int(16*math.Sin(float64(x*2))))
	return color.RGBA{v, v, 255, 255}
}

func exercise09() {
	m := myImage{}
	pic.ShowImage(m)
}
