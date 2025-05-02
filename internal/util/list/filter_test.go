package list

import (
	"github.com/NLCaceres/goth-example/internal/util/test"
	"github.com/google/go-cmp/cmp"
	"strings"
	"testing"
)

func TestFilter(t *testing.T) {
	tests := map[string]struct {
		Input      []string
		Expect     []string
		FilterFunc func(string) bool
	}{
		"Filters out chars after c": {[]string{"c", "d", "e"}, []string{"c"}, func(str string) bool {
			return strings.Compare(str, "c") == 0
		}},
		"Filters with type coercion BUT returns []string": {[]string{"A", "b", "c"}, []string{"A"},
			func(str string) bool { return int(str[0]) < 98 },
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			actual := Filter(testCase.Input, testCase.FilterFunc)
			if !cmp.Equal(testCase.Expect, actual) {
				t.Error(test.ErrorMsg("new list", testCase.Expect, actual))
			}
		})
	}
}

func TestDistinctBy(t *testing.T) {
	tests := map[string]struct {
		Input        []string
		Expect       []string
		SelectorFunc func(string) any
	}{
		"Gets distinct strings by length": {[]string{"a", "ab", "de"}, []string{"a", "ab"},
			func(str string) any { return len(str) },
		},
		"Get unique strings, noting first come first kept": {[]string{"de", "ab"}, []string{"de"},
			func(str string) any { return len(str) },
		},
		"Gets distinct strings ignoring case": {[]string{"AB", "aB", "ab"}, []string{"AB"},
			func(str string) any { return strings.ToLower(str) },
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			actual := DistinctBy(testCase.Input, testCase.SelectorFunc)
			if !cmp.Equal(testCase.Expect, actual) {
				t.Error(test.ErrorMsg("new list", testCase.Expect, actual))
			}
		})
	}
}
