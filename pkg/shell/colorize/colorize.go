package colorize

import "github.com/fatih/color"

func With(attrs ...Attribute) func(format string, a ...any) string {
	return color.New(attrs...).SprintfFunc()
}

func Label(l, format string, args ...any) string {
	return With(Yellow, Bold)("%s: "+format, append([]any{l}, args...)...)
}

// Attribute defines a single SGR Code
type Attribute = color.Attribute

// Base attributes
const (
	Reset        Attribute = color.Reset
	Bold         Attribute = color.Bold
	Faint        Attribute = color.Faint
	Italic       Attribute = color.Italic
	Underline    Attribute = color.Underline
	BlinkSlow    Attribute = color.BlinkSlow
	BlinkRapid   Attribute = color.BlinkRapid
	ReverseVideo Attribute = color.ReverseVideo
	Concealed    Attribute = color.Concealed
	CrossedOut   Attribute = color.CrossedOut
)

// Foreground text colors
const (
	Black   Attribute = color.FgBlack
	Red     Attribute = color.FgRed
	Green   Attribute = color.FgGreen
	Yellow  Attribute = color.FgYellow
	Blue    Attribute = color.FgBlue
	Magenta Attribute = color.FgMagenta
	Cyan    Attribute = color.FgCyan
	White   Attribute = color.FgWhite
)

// Foreground Hi-Intensity text colors
const (
	HiBlack   Attribute = color.FgHiBlack
	HiRed     Attribute = color.FgHiRed
	HiGreen   Attribute = color.FgHiGreen
	HiYellow  Attribute = color.FgHiYellow
	HiBlue    Attribute = color.FgHiBlue
	HiMagenta Attribute = color.FgHiMagenta
	HiCyan    Attribute = color.FgHiCyan
	HiWhite   Attribute = color.FgHiWhite
)

// Background text colors
const (
	BgBlack   Attribute = color.BgBlack
	BgRed     Attribute = color.BgRed
	BgGreen   Attribute = color.BgGreen
	BgYellow  Attribute = color.BgYellow
	BgBlue    Attribute = color.BgBlue
	BgMagenta Attribute = color.BgMagenta
	BgCyan    Attribute = color.BgCyan
	BgWhite   Attribute = color.BgWhite
)

// Background Hi-Intensity text colors
const (
	BgHiBlack   Attribute = color.BgHiBlack
	BgHiRed     Attribute = color.BgHiRed
	BgHiGreen   Attribute = color.BgHiGreen
	BgHiYellow  Attribute = color.BgHiYellow
	BgHiBlue    Attribute = color.BgHiBlue
	BgHiMagenta Attribute = color.BgHiMagenta
	BgHiCyan    Attribute = color.BgHiCyan
	BgHiWhite   Attribute = color.BgHiWhite
)
