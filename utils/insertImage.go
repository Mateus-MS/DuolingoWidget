package utils

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func InsertImage(img *image.RGBA, imagePath string, x, y int) {

	imageFile, err := os.Open(imagePath)
	if err != nil {
		panic(err)
	}
	defer imageFile.Close()

	var imageBits image.Image
	if strings.Contains(imagePath, ".png") {
		imageBits, err = png.Decode(imageFile)
	} else {
		imageBits, err = jpeg.Decode(imageFile)
	}

	if err != nil {
		panic(err)
	}
	dstRect := image.Rect(x, y, x+img.Bounds().Dx(), y+img.Bounds().Dy())

	if strings.Contains(imagePath, ".png") {
		draw.Draw(img, dstRect, imageBits, image.Point{}, draw.Over)
	} else {
		draw.Draw(img, dstRect, imageBits, image.Point{}, draw.Src)
	}

	// draw.Draw(img, img.Bounds(), imageBits, image.Point{}, draw.Src)

}
