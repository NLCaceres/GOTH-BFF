package proxy

import (
	"errors"
	"fmt"
	"github.com/NLCaceres/goth-example/internal/util/test"
	"testing"
)

func TestProxyError(t *testing.T) {
	tests := map[string]struct {
		Input  error
		Msg    string
		Expect bool
	}{
		"RequestError is ProxyError": {
			RequestError{}, "Proxy Error: Proxied request failed for unknown reason", true,
		},
		"RequestError is filled ProxyError": {
			RequestError{"foobar"}, `Proxy Error: Proxied request failed due to "foobar"`, true,
		},
		"Not all errors are ProxyErrors": {errors.New("Foo"), "Proxy Error: Foo", false},
		"Wrapped ProxyError": {
			fmt.Errorf("Some err = %w", RequestError{}),
			"Proxy Error: Some err = Proxied request failed for unknown reason", true,
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if errors.As(testCase.Input, new(Error)) != testCase.Expect {
				t.Error(test.ErrorMsg("Is ProxyError", testCase.Input, testCase.Expect))
			}
			if e, ok := testCase.Input.(Error); ok && e.ProxyError().Error() != testCase.Msg {
				t.Error(test.ErrorMsg("ProxyError msg", testCase.Msg, e.ProxyError().Error()))
			}
		})
	}
}

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
