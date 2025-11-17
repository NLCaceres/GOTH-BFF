package handler

import (
	"github.com/NLCaceres/goth-example/internal/util/test"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestQueryJSON(t *testing.T) {
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
			"internal/test_query.json", "foo|bar|fi", http.StatusOK, `"url":"foo"}]}`,
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
			QueryJSON(c)
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
