package codeview

import (
	"strings"
)

func Language(lang string) func(string) []TextEntry {
	switch lang {
	case "lua":
		return Lua
	case "python":
		return Python
	case "go", "golang":
		return Golang
	default:
		return PlainText
	}
}

func PlainText(line string) []TextEntry {
	return []TextEntry{
		{
			Text:  line,
			Color: ColorDefault,
		},
	}
}

func Lua(line string) []TextEntry {
	tokens := strings.Split(line, " ")

	entries := make([]TextEntry, 0)

	for i, token := range tokens {
		entry := TextEntry{
			Text: token,
		}

		switch token {
		case "local", "function", "end", "if", "then", "else", "elseif", "for", "in", "do", "repeat", "until", "while", "return", "break", "goto", "not", "and", "or":
			entry.Color = ColorToken
		case "true", "false", "nil", "self":
			entry.Color = ColorValue
		}

		if i != 0 {
			entry.Text = " " + entry.Text
		}

		entries = append(entries, entry)
	}

	return entries
}

func Python(line string) []TextEntry {
	tokens := strings.Split(line, " ")

	entries := make([]TextEntry, 0)

	for i, token := range tokens {
		entry := TextEntry{
			Text: token,
		}

		switch token {
		case "def", "class", "return", "if", "elif", "else", "for", "while", "in", "not", "and", "or", "is", "try", "except", "finally", "raise", "assert", "yield", "import", "from", "as", "with", "global", "nonlocal", "del", "pass", "break", "continue", "lambda":
			entry.Color = ColorToken
		case "True", "False", "None":
			entry.Color = ColorValue
		}

		if i != 0 {
			entry.Text = " " + entry.Text
		}

		entries = append(entries, entry)
	}

	return entries
}

func Golang(line string) []TextEntry {
	tokens := strings.Split(line, " ")

	entries := make([]TextEntry, 0)

	for i, token := range tokens {
		entry := TextEntry{
			Text: token,
		}

		switch token {
		case "package", "import", "func", "return", "if", "else", "for", "range", "switch", "case", "default", "type", "var", "const", "struct", "interface", "map", "chan", "go", "select", "break", "continue", "fallthrough", "defer":
			entry.Color = ColorToken
		case "true", "false", "nil":
			entry.Color = ColorValue
		}

		if i != 0 {
			entry.Text = " " + entry.Text
		}

		entries = append(entries, entry)
	}

	return entries
}
