package route

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestMapFromString(t *testing.T) {
	//NOTE: Instead of a []struct slice, using a map catches unexpected data mutations
	var tests = map[string]struct {
		input string
		want  map[string]string
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

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			output := mapFromString(test.input)

			if !cmp.Equal(test.want, output) {
				t.Errorf("Expected %v but got %v", test.want, output)
			}
			for key, value := range output {
				if test.want[key] != value {
					t.Errorf("Key %q should have a value of %q BUT has a value of %q", key, value, test.want[key])
				}
			}
		})
	}
}
