package queryapi

import (
	"github.com/NLCaceres/goth-example/internal/util/test"
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
