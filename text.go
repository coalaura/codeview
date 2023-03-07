package codeview

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"strings"
)

func _lines(text string) []string {
	text = strings.ReplaceAll(text, "\t", "    ")
	text = strings.ReplaceAll(text, "\r\n", "\n")

	lines := strings.Split(text, "\n")

	if len(lines) >= 16 {
		more := len(lines) - 16

		lines = lines[:16]

		lines = append(lines, "... "+_fmtNumber(more)+" more")
	}

	return lines
}

func _fmtNumber(n int) string {
	p := message.NewPrinter(language.English)

	return p.Sprintf("%d", n)
}