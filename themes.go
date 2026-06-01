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
		Background: color.RGBA{R: 36, G: 39, B: 58, A: 255},
		Foreground: color.RGBA{R: 54, G: 58, B: 79, A: 255},

		Default: color.RGBA{R: 202, G: 211, B: 245, A: 255},
		Token:   color.RGBA{R: 198, G: 160, B: 246, A: 255},
		Value:   color.RGBA{R: 238, G: 212, B: 159, A: 255},
		String:  color.RGBA{R: 166, G: 218, B: 149, A: 255},
		Comment: color.RGBA{R: 110, G: 115, B: 141, A: 255},
	}
}
