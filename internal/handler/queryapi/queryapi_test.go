package queryapi

import (
	"errors"
	"github.com/NLCaceres/goth-example/internal/util/fileread"
	"github.com/NLCaceres/goth-example/internal/util/test"
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
		Mock               test.HttpMock
		QueryFile          string
		Filters            string
		ExpectedStatusCode int
		ExpectedResponse   string
	}{ // Probably never will get 501 err from setter while building query
		"Error building query": {httpMock(badData), "./bad.json", "", 500, ""},
		"Error from inside PostJSON": {
			httpMock(badData), "internal/test_query.json", "foo|bar|fi", 502, "",
		},
		"Successfully POSTed to external API": {
			httpMock(successData), "internal/test_query.json", "foo|bar|fi", 200, successData,
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
			Call(c)
			if rec.Code != testCase.ExpectedStatusCode {
				t.Error(test.ErrorMsg("response", testCase.ExpectedStatusCode, rec.Code))
			}
			if strings.TrimSpace(rec.Body.String()) != testCase.ExpectedResponse {
				t.Error(test.ErrorMsg("response", testCase.ExpectedResponse, rec.Body.String()))
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
	return test.HttpMock{RequestMethod: "POST", ResponseStatus: 200, ResponseData: data}
}

func TestQueryErrCode(t *testing.T) {
	tests := map[string]struct {
		Err    error
		Expect int
	}{
		"500 error code if FileReadError": {Err: fileread.FileNotFoundError{}, Expect: 500},
		"501 error code if setter error":  {Err: SearchSetterError, Expect: 501},
		"400 bad request by default":      {Err: errors.New("foo"), Expect: 400},
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
