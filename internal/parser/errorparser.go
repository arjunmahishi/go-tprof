package parser

import (
	"github.com/fatih/color"
	"github.com/gokultp/go-tprof/internal/reports"
)

// ErrorParser parser will parse error messages
type ErrorParser struct {
	text string
}

// NewErrorParser returns a new instance of ErrorParser
func NewErrorParser(line string) *ErrorParser {
	return &ErrorParser{
		text: line,
	}
}

// IsAbleToParse will say this parser is able to parse the given text
func (d *ErrorParser) IsAbleToParse() bool {
	return true
}

// Println will print the line with formatting and colors
func (d *ErrorParser) Println() {
	color.New(color.FgRed).Println(d.text)
	underline := make([]rune, len(d.text))
	for i := 0; i < len(d.text); i++ {
		if d.text[i] == ' ' || d.text[i] == '\t' {
			underline[i] = rune(d.text[i])
			continue
		}
		underline[i] = '^'
	}
	color.New(color.FgRed).Println(string(underline))

}

// UpdateReports will update the reports and temp map by reference
func (d *ErrorParser) UpdateReports(
	r *reports.Report,
	f map[string]*reports.TestFunc,
	failed *string,
) {
	if failed != nil {
		if _, ok := f[*failed]; ok {
			f[*failed].SetError(d.text)
		}
	}

}
