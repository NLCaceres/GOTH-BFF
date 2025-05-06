package stringy

import (
	"github.com/NLCaceres/goth-example/internal/util/test"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"
)

func TestTitleCase(t *testing.T) {
	tests := map[string]struct {
		Input  string
		Expect string
	}{ // Go strings can't be nil so an empty check should suffice
		"An empty string stays empty":     {"", ""},
		"foo becomes Foo":                 {"foo", "Foo"},
		"fOO becomes Foo":                 {"fOO", "Foo"},
		"'foo bar' becomes 'Foo Bar'":     {"foo bar", "Foo Bar"},
		"'Foo Bar' STAYS 'Foo Bar'":       {"Foo Bar", "Foo Bar"},
		"'fIzz buZZ' becomes 'Fizz Buzz'": {"fIzz buZZ", "Fizz Buzz"},
		// Hyphenated words get Titled as individual words
		"'john-smith' becomes 'John-Smith'": {"john-smith", "John-Smith"},
		"'John-Smith' STAYS 'John-Smith'":   {"John-Smith", "John-Smith"},
		// Underscored words get Titled as one word
		"'jack_jill' becomes 'Jack_jill'": {"jack_jill", "Jack_jill"},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			actual := TitleCase(testCase.Input)
			if actual != testCase.Expect {
				t.Error(test.QuotedErrorMsg("TitleCased string", testCase.Expect, actual))
			}
		})
	}
}

func TestFindDunderVars(t *testing.T) {
	tests := map[string]struct {
		Input  string
		Expect []string
	}{
		"No dunder values":         {"foo", []string{}},
		"Incorrect lowercasing":    {"__foo__", []string{}},
		"Incorrect normal casing":  {"__Foo__", []string{}},
		"Missing brackets":         {"__FOO__", []string{"__FOO__"}},
		"Underscored Var name":     {"__DUNDER_VAR__", []string{"__DUNDER_VAR__"}},
		"Backtick word boundary":   {"`__FOO__`", []string{"__FOO__"}},
		"Apostrophe word boundary": {"'__FOO__'", []string{"__FOO__"}},
		"Space word boundary":      {" __FOO__ ", []string{"__FOO__"}},
		"Hyphen word boundary":     {"-__FOO__-", []string{"__FOO__"}},
		"Slash word boundary":      {"/__FOO__/", []string{"__FOO__"}},
		"Star word boundary":       {"***__FOO__***", []string{"__FOO__"}},
		"'()' word boundary":       {"(__FOO__)", []string{"__FOO__"}},
		"'{}' word boundary":       {"{__FOO__}", []string{"__FOO__"}},
		"'<>' word boundary":       {"<__FOO__>", []string{"__FOO__"}},
		"Character word boundary":  {"a__FOO__b", []string{}},
		"Extra underscores":        {"[__FOO___]", []string{}},
		"Extra spaces":             {"[ __FOO__ ]", []string{"__FOO__"}},
		"1 var found":              {"[__FOO__]", []string{"__FOO__"}},
		"2 vars found":             {"[__FOO__][__BAR__]", []string{"__FOO__", "__BAR__"}},
		"Spaced out vars found":    {"[__FOO__] && [__BAR__]", []string{"__FOO__", "__BAR__"}},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			matches, err := FindDunderVars(testCase.Input)
			if err != nil { // Only 1 error possible BUT probably can't manually trigger it
				t.Error("Unexpectedly found a Regexp compilation issue")
			}

			if !cmp.Equal(testCase.Expect, matches, cmpopts.EquateEmpty()) {
				t.Error(test.ErrorMsg("matches", testCase.Expect, matches))
			}
		})
	}
}

func TestMap(t *testing.T) {
	//NOTE: Instead of a []struct slice, using a map catches unexpected data mutations
	tests := map[string]struct {
		Input  string
		Expect map[string]string
	}{ // Laying out maps as follows is probably best for `gofmt`,
		"Empty string": { // otherwise it expects a single line per key-value pair
			"", make(map[string]string),
		},
		"String without comma-separated pairs": {
			"foo", make(map[string]string),
		},
		"String without colon-separated pairs": {
			"fizz,buzz", map[string]string{}, // Equivalent to the make(map) version above
		},
		"String with 1 colon-separated pair": {
			"abc:dec", map[string]string{"abc": "dec"},
		},
		"String with comma-separated pairs": {
			"John:Smith,Jack:Jill", map[string]string{"John": "Smith", "Jack": "Jill"},
		},
		"String with pairs where a key or value is missing": {
			"mike:,:carol", map[string]string{"mike": "", "": "carol"},
		},
		"String with pairs missing a colon": {
			"mar,cia:greg,ja:n,peter,cindy,bobby", map[string]string{"cia": "greg", "ja": "n"},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			output := Map(testCase.Input)

			if !cmp.Equal(testCase.Expect, output) {
				t.Error(test.ErrorMsg("map", testCase.Expect, output))
			}
		})
	}
}

func TestUnescapeUnicodeStr(t *testing.T) {
	tests := map[string]struct {
		Input  string
		Expect string
		Err    string
	}{
		"An empty string returns empty":                  {"", "", ""},
		"Normal strings returns normal":                  {"Foo", "Foo", ""},
		"Normal sentence returns normal":                 {"Foo bar", "Foo bar", ""},
		"String with ASCII":                              {"Foo || Bar", "Foo || Bar", ""},
		"String with ASCII in Sequence":                  {"Foo\u007cBar", "Foo|Bar", ""},
		"String with ASCII in Escaped Sequence":          {"Fi\\u007cZz", "Fi\u007cZz", ""},
		"String with ASCII in Escaped Hex fails":         {"H\\u0x7cI", "", "invalid syntax"},
		"String with ASCII in Unicode":                   {"cómo", "cómo", ""},
		"String with ASCII in Escaped Unicode":           {"c\u00f3mo", "cómo", ""},
		"String with ASCII in Escaped Hex Unicode fails": {"c\\u0xf3mo", "", "invalid syntax"},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			str, err := UnescapeUnicodeStr([]byte(testCase.Input))
			if !test.IsSameError(err, testCase.Err) {
				t.Error(test.QuotedErrorMsg("error", testCase.Err, err))
			}
			if str != testCase.Expect {
				t.Error(test.QuotedErrorMsg("unescaped unicode string", testCase.Expect, str))
			}
		})
	}
}
