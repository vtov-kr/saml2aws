package helper

import (
	"bytes"
	"strings"

	"github.com/fatih/color"
)

// Colorize returns a function that formats a string with the specified ANSI color attributes.
// The returned function behaves like fmt.Sprintf but applies the given color attributes.
func Colorize(attrs ...color.Attribute) func(format string, a ...any) string {
	return color.New(attrs...).SprintfFunc()
}

// ColorizeWithKey returns a function that formats key-value pair and applies ANSI color attributes to the entire output.
func ColorizeWithKey(attrs ...color.Attribute) func(key, format string, args ...any) string {
	return func(key, format string, args ...any) string {
		return Colorize(attrs...)("%s: "+format, append([]any{key}, args...)...)
	}
}

// IndentLines indents each line with the specified number of spaces.
//
// If `prependNewline` is true, a newline is added at the beginning.
// If `appendNewline` is true, a newline is added at the end.
func IndentLines(indent int, prependNewLine, appendNewLine bool, lines ...string) string {
	var (
		nl  byte = '\n'
		pad      = bytes.Repeat([]byte{' '}, indent)
		b   strings.Builder
	)
	if prependNewLine {
		b.WriteByte(nl)
	}
	for i, line := range lines {
		if i > 0 {
			b.WriteByte(nl)
		}
		b.Write(pad)
		b.WriteString(line)
	}
	if appendNewLine {
		b.WriteByte(nl)
	}
	return b.String()
}
