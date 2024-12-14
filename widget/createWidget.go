package widget

import (
	"image"
	"image/color"
	"image/draw"
	"os"
	"path/filepath"

	"github.com/Mateus-MS/DuolingoWidget/utils"
)

func Create(streak string) *image.RGBA {

	img := image.NewRGBA(image.Rect(0, 0, 300, 300))
	bgColor := color.RGBA{R: 255, G: 35, B: 42, A: 0}
	draw.Draw(img, img.Bounds(), &image.Uniform{C: bgColor}, image.Point{}, draw.Src)

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	utils.InsertImage(img, filepath.Join(wd, "assets", "images", "d1.jpg"), 0, 0)
	InsertLabel(img, streak)

	return img

}
