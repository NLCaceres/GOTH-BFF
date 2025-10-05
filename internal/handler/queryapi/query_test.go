package queryapi

import (
	"github.com/NLCaceres/goth-example/internal/util/test"
	"github.com/google/go-cmp/cmp"
	"os"
	"testing"
)

func TestSetFilters(t *testing.T) {
	tests := map[string]struct {
		Input       any
		Replacement string
		Expect      any
		Err         string
	}{
		"No matches found": {Input: "foo", Expect: "foo"},
		"One match found but multiple replacements": { // NEED ALL CAPS DunderVars
			Input: "[__FOO__]", Replacement: "foo|bar", Expect: "[foo]",
		},
		"One replacement but multiple matches": {
			Input: "[__FOO__] && [__BAR__]", Replacement: "fi", Expect: "[fi] && [__BAR__]",
		},
		"All replacements successful": {
			Input: "[__FOO__] && [__BAR__]", Replacement: "foo|bar", Expect: "[foo] && [bar]",
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			search := Search{FilterBy: testCase.Input.(string)}
			os.Setenv("FILTER_REPLACEMENTS", testCase.Replacement)
			err := search.setFilters(testCase.Replacement)
			if search.FilterBy != testCase.Expect {
				t.Error(test.ErrorMsg("filter", testCase.Expect, search.FilterBy))
			}
			if !test.IsSameError(err, testCase.Err) {
				t.Error(test.QuotedErrorMsg("error", testCase.Err, err))
			}
		})
	}
}

func TestRequestLast(t *testing.T) {
	tests := map[string]struct {
		Input  Request
		Expect *Search
	}{
		"Return nil if 0 Terms":   {Request{[]Search{}}, nil},
		"Return only Search Term": {Request{[]Search{{Q: "Foo"}}}, &Search{Q: "Foo"}},
		"Return last Search Term": {Request{[]Search{{Q: "F"}, {Q: "B"}}}, &Search{Q: "B"}},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if !cmp.Equal(testCase.Expect, testCase.Input.last()) {
				t.Error(test.ErrorMsg("last search term", testCase.Expect, testCase.Input.last()))
			}
		})
	}
}
