package util

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TitleCase(someString string) string {
	return cases.Title(language.English).String(someString)
}
