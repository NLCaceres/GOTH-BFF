package queryapi

import (
	"errors"
	"github.com/NLCaceres/goth-example/internal/util/fileread"
	"github.com/NLCaceres/goth-example/internal/util/test"
	"github.com/google/go-cmp/cmp"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestCall(t *testing.T) {
	badData := `"foo":"bar"`
	successData := "{" + badData + "}"
	tests := map[string]struct {
		Mock      test.HttpMock
		QueryFile string
		Filters   string
		Expect    map[string]any
		Err       any
	}{ // Probably never will get 501 err from setter while building query
		"Error building query": {
			httpMock(badData), "./bad.json", "", nil, new(fileread.FileReadError),
		},
		"Error from inside PostJSON": {
			httpMock(badData), "internal/test_query.json", "foo|bar|fi", nil, new(error),
		},
		"Successfully POSTed to external API": {
			httpMock(successData), "internal/test_query.json",
			"foo|bar|fi", map[string]any{"foo": "bar"}, nil,
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			server := httptest.NewServer(test.HttpHandlerFunc(t, testCase.Mock))
			defer server.Close()
			os.Setenv("EXTERNAL_API_URL", server.URL)
			// Actual request from client NOT relevant to the test since only grabbing its Path
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
			rec := httptest.NewRecorder()
			c := echo.New().NewContext(req, rec)
			c.SetPath("/foo") //NOTE: BUT must set the Path in Echo's context here!

			os.Setenv("QUERY_FILE", testCase.QueryFile)
			os.Setenv("FILTER_REPLACEMENTS", testCase.Filters)
			res, err := Call(c)
			nilErr := testCase.Err == nil
			if (nilErr && err != nil) || (!nilErr && !errors.As(err, &testCase.Err)) {
				t.Error(test.ErrorMsg("error", testCase.Err, err))
			}
			if !cmp.Equal(res, testCase.Expect) {
				t.Error(test.ErrorMsg("response", testCase.Expect, res))
			}
		})
	}
}

func TestNewQuery(t *testing.T) {
	tests := map[string]struct {
		Input     string
		QueryFile string
		Err       any
		Expect    string
	}{
		"Returns ready query":   {"foo", "internal/test_query.json", nil, "foo"},
		"Empty query is empty":  {"", "internal/test_query.json", nil, ""},
		"File not found error":  {"", "./bad.json", fileread.FileNotFoundError{}, ""},
		"File formatting error": {"", "./bad_typing.json", fileread.MalformedJsonError{}, ""},
	} // Unsure how to get `Marshal` to fail so no test case for it
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			os.Setenv("QUERY_FILE", testCase.QueryFile)
			query, err := newQuery(testCase.Input)
			nilErr := testCase.Err == nil
			if (!nilErr && !errors.As(err, &testCase.Err)) || (nilErr && err != nil) {
				t.Error(test.ErrorMsg("error", testCase.Err, err))
			}
			expect := `"q": "` + testCase.Expect + `",`
			if (nilErr && !strings.Contains(query.String(), expect)) || (!nilErr && query != nil) {
				t.Error(test.ErrorMsg("Path input", testCase.Input, query))
			}
		})
	}
}

func httpMock(data string) test.HttpMock {
	return test.HttpMock{RequestMethod: "POST", ResponseStatus: http.StatusOK, ResponseData: data}
}

func TestQueryErrCode(t *testing.T) {
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
			actual := queryErrCode(testCase.Err)
			if actual != testCase.Expect {
				t.Error(test.ErrorMsg("Error code", testCase.Expect, actual))
			}
		})
	}
}
