package main

import(
	"strings"
)


// Credit: https://stackoverflow.com/a/54426140
func SplitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}