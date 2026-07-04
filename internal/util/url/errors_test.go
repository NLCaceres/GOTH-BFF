package url

import (
	"errors"
	"fmt"
	"github.com/NLCaceres/goth-example/internal/util/test"
	"testing"
)

func TestUrlError(t *testing.T) {
	tests := map[string]struct {
		Input  error
		Msg    string
		Expect bool
	}{
		"ParamError is UrlError": {
			ParamError{}, "Url Error: URL parameter invalid BUT unknown key and value", true,
		},
		"ParamError filled is UrlError": {
			ParamError{"a", "b"}, "Url Error: Invalid URL parameter: a - b", true,
		},
		"Not all errors are UrlErrors": {errors.New("foo"), "URL Error: foo", false},
		"Wrapped URLError": {
			fmt.Errorf("Some err = %w", ParamError{}),
			"Url Error: URL parameter invalid BUT unknown key and value", true,
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			var e Error
			if errors.As(testCase.Input, &e) != testCase.Expect {
				t.Error(test.ErrorMsg("`As` UrlError", testCase.Expect, !testCase.Expect))
			}
			if e != nil && e.UrlError().Error() != testCase.Msg {
				t.Error(test.ErrorMsg("UrlError msg", testCase.Msg, e.UrlError().Error()))
			}
		})
	}

}

func TestParamError(t *testing.T) {
	tests := map[string]struct {
		Key    string
		Value  string
		Expect string
	}{
		"Missing key & value": {"", "", "URL parameter invalid BUT unknown key and value"},
		"Blank key":           {"   ", "f", "URL parameter invalid BUT value=f with unknown key"},
		"Blank value":         {"k", "  ", "URL parameter invalid with key=k and unknown value"},
		"With key & value":    {"ab", "cd", "Invalid URL parameter: ab - cd"},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			e := ParamError{Key: testCase.Key, Value: testCase.Value}
			if testCase.Expect != e.Error() {
				t.Error(test.ErrorMsg("Param error", testCase.Expect, e.Error()))
			}
		})
	}
}
