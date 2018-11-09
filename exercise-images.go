package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	// ColorModel returns the Image's color model.
	color.Model
	image.Rectangle
	color.Color
}

func main() {
	m := Image{color.RGBAModel, image.Rect(0, 0, 100, 100), color.RGBA{255, 255, 255, 255}}
	pic.ShowImage(m)
}
