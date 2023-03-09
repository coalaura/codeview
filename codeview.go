package codeview

import (
	"bytes"
	_ "embed"
	"github.com/golang/freetype/truetype"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	"image"
	"image/color"
	"image/png"
	"math"
)

var (
	//go:embed assets/font.ttf
	fontBytes []byte

	logo        []byte = nil
	projectName        = ""
)

// SetProjectName sets the project name to be used.
func SetProjectName(name string) {
	projectName = name
}

// SetLogo sets the logo image to be used.
func SetLogo(png []byte) {
	logo = png
}

func _render(text, title string, scale float64) (image.Image, error) {
	width := math.Floor(802 * scale)
	height := math.Floor(528 * scale)

	dest := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))

	background := color.RGBA{R: 30, G: 30, B: 46, A: 255}
	foreground := color.RGBA{R: 38, G: 38, B: 55, A: 255}
	textColor := color.RGBA{R: 173, G: 186, B: 199, A: 255}
	mutedColor := color.RGBA{R: 70, G: 77, B: 83, A: 255}

	pad := 25 * scale
	fontSize := 18 * scale
	lineHeight := 28 * scale
	logoY := 22 * scale

	innerX := pad + (18 * scale)
	innerY := innerX + fontSize + (60 * scale)

	gc := draw2dimg.NewGraphicContext(dest)

	gc.SetFillColor(background)
	draw2dkit.Rectangle(gc, 0, 0, width, height)

	gc.Fill()

	gc.SetFillColor(foreground)
	draw2dkit.Rectangle(gc, pad, pad+(60*scale), width-pad, height-pad)

	gc.Fill()

	logoImg, err := png.Decode(bytes.NewReader(logo))
	if err != nil {
		return nil, err
	}

	gc.Save()

	w := float64(logoImg.Bounds().Dx())
	sc := (48 / w) * scale

	gc.Translate(pad, logoY)
	gc.Scale(sc, sc)
	gc.DrawImage(logoImg)

	gc.Restore()

	font, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}

	fontData := draw2d.FontData{
		Name:   "ComicCode",
		Family: draw2d.FontFamilyMono,
		Style:  draw2d.FontStyleNormal,
	}

	draw2d.RegisterFont(fontData, font)

	gc.SetFontData(fontData)

	gc.SetFillColor(textColor)
	gc.SetFontSize(24 * scale)

	gc.FillStringAt(title, pad+(48+16)*scale, logoY+(38*scale))

	gc.SetFontSize(fontSize)

	lines := _lines(text)

	for i, line := range lines {
		y := innerY + float64(i)*lineHeight

		if len(line) > 65 {
			line = line[:65]
		}

		if i == 13 {
			gc.SetFillColor(mutedColor)
		}

		gc.FillStringAt(line, innerX, y)
	}

	return dest, nil
}
