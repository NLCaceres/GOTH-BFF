package util

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func TitleCase(someString string) string {
	return cases.Title(language.English).String(someString)
}

func FindDunderVars(str string) ([]string, error) {
	// Regex to find "__FOO__" text groups surrounded by any word boundary (![A-Za-z0-9_])
	re, err := regexp.Compile(`\b_{2}[A-Z]+_{2}\b`)
	if err != nil {
		log.Print("Issue with regex to find search filters:", err)
		return []string{}, err
	}

	return re.FindAllString(str, -1), nil
}

// Converts a string containing Unicode representations of characters
// like "&" or accented letters like in "está" into readable when printed strings.
// Ex: "c\u00f3mo est\u00e1s" -> "cómo estás"
// Commonly occurs due to Go's `Unmarshal` & `Marshal` escaping these characters
// into their Unicode sequence representations, i.e. "\u0026" or "\u00e1"
func UnescapeUnicodeStr(jsonRaw []byte) (string, error) {
	// MUST Quote, Replace Unicode (\uXXXX) chars, & then unquote/escape to render Unicode
	// JUST Unquoting won't work; It'll return an empty string + err
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(jsonRaw)), `\\u`, `\u`, -1))
	if err != nil {
		return "", err
	}
	return str, nil
}
