package url

import (
	"github.com/NLCaceres/goth-example/internal/util/test"
	"net/url"
	"testing"
)

func TestIntQuery(t *testing.T) {
	tests := map[string]struct {
		Input  string
		Expect int
		Err    error
	}{
		"Missing param": {"", 1, nil},
		"Empty param":   {"foo", 0, ParamError{Key: "foo"}},
		// Browser would trim the param value so not usually get a blank BUT an empty instead
		"Blank param":   {"foo=    ", 0, ParamError{Key: "foo", Value: "    "}},
		"Invalid param": {"foo=a", 0, ParamError{Key: "foo", Value: "a"}},
		"Valid param":   {"foo=123", 123, nil},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			u := URL{url.URL{RawQuery: testCase.Input}}
			q, err := u.QueryInt("foo")
			if err != testCase.Err { // Could use `errors.Is` BUT expecting a ParamErr exactly
				t.Error(test.ErrorMsg("Int Query Param err", testCase.Err, err))
			}
			if testCase.Expect != q {
				t.Error(test.QuotedErrorMsg("Int Query Param", testCase.Expect, q))
			}
		})
	}
}
