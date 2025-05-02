package list

import (
	"github.com/NLCaceres/goth-example/internal/util/test"
	"github.com/google/go-cmp/cmp"
	"strconv"
	"testing"
)

func TestForEach(t *testing.T) {
	mapper := func(num string) (int, error) { return strconv.Atoi(num) }
	tests := map[string]struct {
		Input  []string
		Expect []int
		Err    string
	}{
		"Coerces string slice into int slice": {[]string{"1", "2", "3"}, []int{1, 2, 3}, ""},
		"Mapper fails":                        {[]string{"1", "b", "c"}, []int{}, "invalid syntax"},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			actual, err := ForEach(testCase.Input, mapper)
			if !cmp.Equal(testCase.Expect, actual) {
				t.Error(test.ErrorMsg("new list", testCase.Expect, actual))
			}
			if !test.IsSameError(err, testCase.Err) {
				t.Error(test.QuotedErrorMsg("error", testCase.Err, err))
			}
		})
	}
}
