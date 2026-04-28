package internal

import "strings"

func Blank(s string) string {
	if len(s) == 0 {
		return `""`
	}
	return s
}

func ShowNewlines(s string) string {
	var buf strings.Builder
	for _, c := range s {
		if c == '\n' {
			buf.WriteRune('\u2424') // ␤
		}
		buf.WriteRune(c)
	}
	return buf.String()
}
