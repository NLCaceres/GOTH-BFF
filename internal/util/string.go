package util

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"regexp"
	"strings"
)

func TitleCase(someString string) string {
	return cases.Title(language.English).String(someString)
}

func FindDunderVars(str string) ([]string, error) {
	// Regex to find "[__FOO__]" text groups. All else fails
	re, err := regexp.Compile(`\[_{2}[A-Z]+_{2}\]`) // MUST Trim remaining "[" later
	if err != nil {
		log.Print("Issue with regex to find search filters:", err)
		return []string{}, err
	}

	matches := re.FindAllString(str, -1)
	trimmedMatches := make([]string, len(matches))
	for i, match := range matches { // Remove surrounding brackets
		trimmedMatches[i] = strings.Trim(match, "[]")
	}
	return trimmedMatches, nil
}
