package stringy

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
	re, err := regexp.Compile(`\b_{2}[A-Z]+(_?[A-Z])*_{2}\b`)
	if err != nil {
		log.Print("Issue with regex to find search filters:", err)
		return []string{}, err
	}

	return re.FindAllString(str, -1), nil
}

// Converts a string with pairs separated by commas and pair key separated from its
// value by colons into a map[string]string.
//
// Ex: "foo:bar,fizz:buzz" -> map[string]string{"foo":"bar", "fizz":"buzz"}
func Map(formattedString string) map[string]string {
	newMap := make(map[string]string)
	//NOTE: Golang ISN'T functional so no `map`, `filter`, etc. There's only `for`
	for _, keyValPair := range strings.Split(formattedString, ",") {
		splitKeyVal := strings.Split(keyValPair, ":") // [key, value]
		if len(splitKeyVal) > 1 {
			key, value := splitKeyVal[0], splitKeyVal[1]
			newMap[key] = value
		}
	}
	return newMap
}

// Converts a string containing Unicode representations of characters
// like "&" or accented letters like in "está" into readable when printed strings.
//
// Ex: "c\u00f3mo est\u00e1s" -> "cómo estás"
//
// Commonly occurs due to Go's `Unmarshal` & `Marshal` escaping these characters
// into their Unicode sequence representations, i.e. "\u0026" or "\u00e1"
func UnescapedUnicode(jsonRaw []byte) (string, error) {
	// MUST Quote, Replace Unicode (\uXXXX) chars, & then unquote/escape to render Unicode
	// JUST Unquoting won't work; It'll return an empty string + err
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(jsonRaw)), `\\u`, `\u`, -1))
	if err != nil {
		return "", err
	}
	return str, nil
}
