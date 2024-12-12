package widget

import (
	"image"
	"image/color"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

func InsertLabel(img *image.RGBA, streak string) {

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fontBytes, err := os.ReadFile(filepath.Join(wd, "assets", "fonts", "MADE_Tommy_Soft_Bold_PERSONAL_USE.otf")) // Change to your font's path
	if err != nil {
		log.Fatal(err)
	}

	f, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatalf("failed to parse font: %v", err)
	}
	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    60,
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
	textHeight := face.Metrics().Height.Round()

	centerX := (img.Bounds().Dx() - textWidth) / 2
	centerY := textHeight - 5

	d.Dot = fixed.Point26_6{fixed.Int26_6(centerX) << 6, fixed.Int26_6(centerY) << 6}
	d.DrawString(streak)
}
