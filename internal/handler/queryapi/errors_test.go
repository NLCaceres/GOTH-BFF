package queryapi

import (
	"errors"
	"github.com/NLCaceres/goth-example/internal/util/fileread"
	"github.com/NLCaceres/goth-example/internal/util/proxy"
	"github.com/NLCaceres/goth-example/internal/util/test"
	"net/http"
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

func TestNewError(t *testing.T) {
	tests := map[string]struct {
		Input  error
		Expect Error
	}{
		"Proxy Err": {
			proxy.RequestError{}, Error{proxy.RequestError{}, http.StatusBadGateway},
		},
		"FileRead Error": {
			fileread.FileNotFoundError{}, Error{fileread.FileNotFoundError{}, http.StatusInternalServerError},
		},
		"Setter Error": {
			SearchSetterError, Error{SearchSetterError, http.StatusNotImplemented},
		},
		"Any error": {
			errors.New("foo"), Error{errors.New("foo"), http.StatusBadRequest},
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			err := NewError(testCase.Input)
			if err.Error() != testCase.Expect.Error() || err.Code != testCase.Expect.Code {
				t.Error(test.QuotedErrorMsg("New QueryAPI Error", testCase.Expect, err))
			}
		})
	}
}

func TestErrCode(t *testing.T) {
	tests := map[string]struct {
		Err    error
		Expect int
	}{
		"500 error code if FileReadError": {
			Err: fileread.FileNotFoundError{}, Expect: http.StatusInternalServerError,
		},
		"501 error code if setter error": {
			Err: SearchSetterError, Expect: http.StatusNotImplemented,
		},
		"400 bad request by default": {
			Err: errors.New("foo"), Expect: http.StatusBadRequest,
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			actual := ErrCode(testCase.Err)
			if actual != testCase.Expect {
				t.Error(test.ErrorMsg("Error code", testCase.Expect, actual))
			}
		})
	}
}
