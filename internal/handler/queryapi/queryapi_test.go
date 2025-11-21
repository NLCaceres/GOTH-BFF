package queryapi

import (
	"github.com/NLCaceres/goth-example/internal/util/fileread"
	"github.com/NLCaceres/goth-example/internal/util/test"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestCall(t *testing.T) {
	tests := map[string]struct {
		Mock      test.HttpMock
		QueryFile string
		Filters   string
		Expect    *Response
		Err       any
	}{ // Probably never will get 501 err from setter while building query
		"Error building query": {
			httpMock(`"foo":"bar"`), "./bad.json", "", nil, new(fileread.Error),
		},
		"Error from inside PostJSON": {
			httpMock(`"foo":"bar"`), "internal/test_query.json", "foo|bar|fi", nil, new(error),
		},
		"Successfully POSTed to external API": {
			httpMock(`{"results": [{ "hits": [{ "document": {"url": "foo"} }] }]}`),
			"internal/test_query.json", "foo|bar|fi", NewResponse([]Document{{URL: "foo"}}), nil,
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
			res, err := Call(c.Path()[1:])
			if !test.EqualErrors(err, testCase.Err) {
				t.Errorf("Expected error = %T but got %#v", testCase.Err, err)
			}
			if !cmp.Equal(res, testCase.Expect, cmpopts.IgnoreUnexported(Response{})) {
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
		"Returns ready query":  {"foo", "internal/test_query.json", nil, "foo"},
		"Empty query is empty": {"", "internal/test_query.json", nil, ""},
		"File not found error": {"", "bad.json", new(fileread.FileNotFoundError), ""},
		"File formatting error": {
			"", "internal/util/test/bad_typing.json", new(fileread.MalformedJsonError), "",
		},
	} // Unsure how to get `Marshal` to fail so no test case for it
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			os.Setenv("QUERY_FILE", testCase.QueryFile)
			query, err := newQuery(testCase.Input)
			if !test.EqualErrors(err, testCase.Err) {
				t.Errorf("Expected error = %T but got %#v", testCase.Err, err)
			}
			expect := `"q": "` + testCase.Expect + `",`
			if query.String() != "<nil>" && !strings.Contains(query.String(), expect) {
				t.Error(test.ErrorMsg("Path input", testCase.Input, query))
			}
		})
	}
}

func httpMock(data string) test.HttpMock {
	return test.HttpMock{RequestMethod: "POST", ResponseStatus: http.StatusOK, ResponseData: data}
}
