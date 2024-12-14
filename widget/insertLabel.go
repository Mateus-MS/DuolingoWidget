package widget

import (
	"image"
	"image/color"
	"log"
	"os"
	"path/filepath"

	"github.com/Mateus-MS/DuolingoWidget/utils"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

func InsertLabel(img *image.RGBA, streak string) {

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fontBytes, err := os.ReadFile(filepath.Join(wd, "assets", "fonts", "MADE_Tommy_Soft_Bold_PERSONAL_USE.otf"))
	if err != nil {
		log.Fatal(err)
	}

	f, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatalf("failed to parse font: %v", err)
	}
	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    70,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	if err != nil {
		log.Fatalf("failed to create new face: %v", err)
	}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.RGBA{255, 255, 255, 255}),
		Face: face,
	}

	textWidth := d.MeasureString(streak).Round()

	streakRightPadding := 8

	totalWidth := 47 + streakRightPadding + textWidth
	padding := (img.Bounds().Dx() - totalWidth) / 2

	utils.InsertImage(img, filepath.Join(wd, "assets", "images", "s.png"), padding, 27)

	d.Dot = fixed.Point26_6{fixed.Int26_6(padding+47+streakRightPadding) << 6, fixed.Int26_6(81) << 6}
	d.DrawString(streak)

}
