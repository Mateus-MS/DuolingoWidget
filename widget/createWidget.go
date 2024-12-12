package widget

import (
	"image"
	"image/color"
	"image/draw"
	"os"
	"path/filepath"
)

func Create(streak string) *image.RGBA {

	img := image.NewRGBA(image.Rect(0, 0, 200, 200))
	bgColor := color.RGBA{R: 255, G: 35, B: 42, A: 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{C: bgColor}, image.Point{}, draw.Src)

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	InsertImage(img, filepath.Join(wd, "assets", "images", "model_01.jpg"))
	InsertLabel(img, streak)

	return img

}
