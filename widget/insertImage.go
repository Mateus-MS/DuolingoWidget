package widget

import (
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

func InsertImage(img *image.RGBA, imagePath string) {

	bgFile, err := os.Open(imagePath)
	if err != nil {
		panic(err)
	}
	defer bgFile.Close()

	bgImage, err := jpeg.Decode(bgFile)
	if err != nil {
		panic(err)
	}

	draw.Draw(img, img.Bounds(), bgImage, image.Point{}, draw.Src)

}
