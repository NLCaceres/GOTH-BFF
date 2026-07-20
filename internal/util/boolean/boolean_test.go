package boolean

import (
	"fmt"
	"github.com/NLCaceres/goth-example/internal/util/test"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestTernary(t *testing.T) {
	tests := map[string]struct {
		Values    []any
		Condition bool
		Expect    any
	}{
		"True value":     {[]any{"12", 12}, true, "12"}, // Types don't actually need to match
		"False value":    {[]any{"foo", 1.2}, false, 1.2},
		"Strings":        {[]any{"a", "b"}, true, "a"}, // BUT gets an `any` type
		"Ints":           {[]any{123, 987}, false, 987},
		"Int comparison": {[]any{123, 987}, 987 > 123, 123},
		"Array":          {[]any{[]string{"a", "b"}, []int{1, 2}}, true, []string{"a", "b"}},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			actual := Ternary(testCase.Condition, testCase.Values[0], testCase.Values[1])
			if !cmp.Equal(testCase.Expect, actual) {
				t.Error(test.QuotedErrorMsg(
					fmt.Sprintf("Ternary %t value", testCase.Condition), testCase.Expect, actual),
				)
			}
		})
	}
}
