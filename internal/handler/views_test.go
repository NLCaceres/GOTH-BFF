package handler

import (
	"errors"
	"github.com/NLCaceres/goth-example/internal/handler/queryapi"
	"github.com/NLCaceres/goth-example/internal/util/fileread"
	"github.com/NLCaceres/goth-example/internal/util/test"
	"net/http"
	"testing"
)

func TestQueryErrCode(t *testing.T) {
	tests := map[string]struct {
		Err    error
		Expect int
	}{
		"500 error code if FileReadError": {
			Err: fileread.FileNotFoundError{}, Expect: http.StatusInternalServerError,
		},
		"501 error code if setter error": {
			Err: queryapi.SearchSetterError, Expect: http.StatusNotImplemented,
		},
		"400 bad request by default": {
			Err: errors.New("foo"), Expect: http.StatusBadRequest,
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			actual := queryErrCode(testCase.Err)
			if actual != testCase.Expect {
				t.Error(test.ErrorMsg("Error code", testCase.Expect, actual))
			}
		})
	}
}
