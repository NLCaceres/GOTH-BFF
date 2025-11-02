package handler

import (
	"errors"
	"github.com/NLCaceres/goth-example/internal/handler/queryapi"
	"github.com/NLCaceres/goth-example/internal/util/fileread"
	"github.com/NLCaceres/goth-example/internal/util/test"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestInlineQueries(t *testing.T) {
	tests := map[string]struct {
		Mock               test.HttpMock
		QueryFile          string
		Filters            string
		ExpectedStatusCode int
		ExpectedResponse   string
	}{ // Probably never will get 501 err from setter while building query
		"Error building query": {
			httpMock(`"foo":"bar"`), "./bad.json", "", http.StatusInternalServerError, "",
		},
		"Error from inside PostJSON": {
			httpMock(`"foo":"bar"`), "internal/test_query.json", "foo|bar|fi", http.StatusBadGateway, "",
		},
		"Successfully POSTed to external API": {
			httpMock(`{"results": [{ "hits": [{ "document": {"url": "foo"} }] }]}`),
			"internal/test_query.json", "foo|bar|fi", http.StatusOK, `"url":"foo"`,
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
			InlineQueries(c)
			if rec.Code != testCase.ExpectedStatusCode {
				t.Error(test.ErrorMsg("response", testCase.ExpectedStatusCode, rec.Code))
			}
			notEquallyEmpty := rec.Body.String() != "" && testCase.ExpectedResponse == ""
			missingQ := !strings.Contains(rec.Body.String(), testCase.ExpectedResponse)
			if notEquallyEmpty || missingQ {
				t.Error(test.ErrorMsg("response", testCase.ExpectedResponse, rec.Body.String()))
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
