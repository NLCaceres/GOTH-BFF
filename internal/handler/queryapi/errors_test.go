package queryapi

import (
	"errors"
	"github.com/NLCaceres/goth-example/internal/util/test"
	"testing"
)

func TestQueryApiError(t *testing.T) {
	tests := map[string]struct {
		Input  Error
		Expect string
	}{
		"Appends cause string to end of 'Queried api: ' prefix": {
			Error{errors.New("foo"), 0}, `Queried api: "foo"`,
		},
		"Nil 'Err' field": {
			Error{}, `Queried api: Unknown error`,
		},
		"Whitespaced 'Err' field message": {
			Error{errors.New("   "), 0}, `Queried api: "   "`,
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if testCase.Input.Error() != testCase.Expect {
				t.Error(test.ErrorMsg(
					"QueryAPI Error message", testCase.Expect, testCase.Input.Error(),
				))
			}
		})
	}
}
