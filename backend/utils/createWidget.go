package utils

import (
	"image"
	"image/color"
	"image/draw"
)

func CreateWidget(days string) *image.RGBA {

	img := image.NewRGBA(image.Rect(0, 0, 500, 200))
	bgColor := color.RGBA{R: 32, G: 35, B: 42, A: 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{C: bgColor}, image.Point{}, draw.Src)

	InsertLabel(img, 50, 100, "Streak: "+days, color.RGBA{R: 97, G: 218, B: 251, A: 255})

	return img

}
