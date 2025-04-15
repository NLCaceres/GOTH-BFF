package strings

import (
	"strings"
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
			if TitleCase(testCase.Input) != testCase.Expect {
				t.Errorf("%q did not become %q as expected", testCase.Input, testCase.Expect)
			}
		})
	}
}

func TestFindDunderVars(t *testing.T) {
	tests := map[string]struct {
		input  string
		expect []string
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
			matches, err := FindDunderVars(testCase.input)
			if err != nil { // Only error possible BUT probably can't manually trigger it
				t.Error("Unexpectedly found a Regexp compilation issue")
			}
			for i, match := range matches {
				if match != testCase.expect[i] {
					t.Errorf("Matches found = %v BUT expected = %v", matches, testCase.expect)

				}
			}
		})
	}
}

func TestUnescapeUnicodeStr(t *testing.T) {
	tests := map[string]struct {
		Input    string
		Expected string
		Err      string
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
			if err != nil && (testCase.Err == "" || !strings.HasPrefix(err.Error(), testCase.Err)) {
				t.Errorf("Expected err = %q BUT got %q", testCase.Err, err)
			}
			if str != testCase.Expected {
				t.Errorf("Expected %q but got %q", testCase.Expected, str)
			}
		})
	}
}
