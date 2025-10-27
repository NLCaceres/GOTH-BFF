package proxy

import (
	"github.com/NLCaceres/goth-example/internal/util/test"
	"testing"
)

func TestRequestError(t *testing.T) {
	tests := map[string]struct {
		Input  RequestError
		Expect string
	}{
		"Appends cause string to end of 'Proxied request failed due to ' prefix": {
			RequestError{"foobar"}, `Proxied request failed due to "foobar"`,
		},
		"Blank 'Cause' field message": {
			RequestError{}, "Proxied request failed for unknown reason",
		},
		"Whitespaced 'Cause' field message": {
			RequestError{"    "}, "Proxied request failed for unknown reason",
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if testCase.Input.Error() != testCase.Expect {
				t.Error(test.ErrorMsg(
					"Error message", testCase.Expect, testCase.Input.Error(),
				))
			}
		})
	}
}
