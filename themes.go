package codeview

import "image/color"

const (
	ColorDefault = iota
	ColorToken
	ColorValue
	ColorString
	ColorComment
)

type Theme struct {
	Background color.RGBA
	Foreground color.RGBA

	Default color.RGBA
	Token   color.RGBA
	Value   color.RGBA
	String  color.RGBA
	Comment color.RGBA
}

func DefaultTheme() Theme {
	return Theme{
		Background: color.RGBA{R: 30, G: 30, B: 46, A: 255},
		Foreground: color.RGBA{R: 38, G: 38, B: 55, A: 255},

		Default: color.RGBA{R: 173, G: 186, B: 199, A: 255},
		Token:   color.RGBA{R: 180, G: 94, B: 164, A: 255},
		Value:   color.RGBA{R: 231, G: 206, B: 86, A: 255},
		String:  color.RGBA{R: 79, G: 180, B: 215, A: 255},
		Comment: color.RGBA{R: 70, G: 77, B: 83, A: 255},
	}
}
