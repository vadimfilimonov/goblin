package goblin

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// UpperFirst converts the first character of `string` to upper case.
func UpperFirst(text string) string {
	return cases.Title(language.Und, cases.NoLower).String(text)
}
