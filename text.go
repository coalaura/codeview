package codeview

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"image/color"
	"strings"
)

type TextEntry struct {
	Text  string
	Color int
}

type Text struct {
	Lines [][]TextEntry
}

func NewText(text string, lineParser func(string) []TextEntry) Text {
	lines, more := _splitText(text)

	t := Text{
		Lines: make([][]TextEntry, len(lines)),
	}

	for i, line := range lines {
		parsed := lineParser(line)
		cleaned := make([]TextEntry, 0)

		length := 0

		for _, entry := range parsed {
			l := len(entry.Text)

			if length+l > 50 {
				l = 50 - length

				entry.Text = entry.Text[:l]

				cleaned = append(cleaned, entry)

				break
			}

			cleaned = append(cleaned, entry)

			length += l
		}

		t.Lines[i] = cleaned
	}

	if more > 0 {
		t.Lines = append(t.Lines, []TextEntry{
			{
				Text:  "... " + _fmtNumber(more) + " more",
				Color: ColorComment,
			},
		})
	}

	return t
}

func (e TextEntry) GetColor(theme Theme) color.RGBA {
	switch e.Color {
	case ColorDefault:
		return theme.Default
	case ColorToken:
		return theme.Token
	case ColorValue:
		return theme.Value
	case ColorString:
		return theme.String
	case ColorComment:
		return theme.Comment
	default:
		return theme.Default
	}
}

func _splitText(text string) ([]string, int) {
	text = strings.ReplaceAll(text, "\t", "    ")
	text = strings.ReplaceAll(text, "\r\n", "\n")

	lines := strings.Split(text, "\n")
	more := 0

	if len(lines) > 14 {
		more = len(lines) - 13

		lines = lines[:13]
	}

	return lines, more
}

func _fmtNumber(n int) string {
	p := message.NewPrinter(language.English)

	return p.Sprintf("%d", n)
}
