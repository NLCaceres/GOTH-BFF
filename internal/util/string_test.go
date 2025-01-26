package util

import "testing"

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
		"No dunder values":      {"foo", []string{}},
		"Missing brackets":      {"__FOO__", []string{}},
		"Extra underscores":     {"[__FOO___]", []string{}},
		"Extra spaces":          {"[ __FOO__ ]", []string{}},
		"1 var found":           {"[__FOO__]", []string{"__FOO__"}},
		"2 vars found":          {"[__FOO__][__BAR__]", []string{"__FOO__", "__BAR__"}},
		"Spaced out vars found": {"[__FOO__] && [__BAR__]", []string{"__FOO__", "__BAR__"}},
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
