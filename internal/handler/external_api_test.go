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

func TestApiPostRequest(t *testing.T) {
	badData := `"foo":"bar"`
	successData := "{" + badData + "}"
	tests := map[string]struct {
		Mock               test.HttpMock
		QueryFile          string
		Filters            string
		ExpectedStatusCode int
		ExpectedResponse   string
	}{
		"Error from inside fileread.JSON()": {httpMock(badData), "./bad.json", "", 500, ""},
		"Error setting filters": {
			httpMock(badData), "internal/util/test/bad_typing.json", "", 501, "",
		},
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
			ApiPostRequest(c)
			if rec.Code != testCase.ExpectedStatusCode {
				t.Error(test.ErrorMsg("response", testCase.ExpectedStatusCode, rec.Code))
			}
			if strings.TrimSpace(rec.Body.String()) != testCase.ExpectedResponse {
				t.Error(test.ErrorMsg("response", testCase.ExpectedResponse, rec.Body.String()))
			}
		})
	}
}

func httpMock(data string) test.HttpMock {
	return test.HttpMock{RequestMethod: "POST", ResponseStatus: 200, ResponseData: data}
}

func TestSetFilters(t *testing.T) {
	tests := map[string]struct {
		Input       any
		Replacement string
		Expect      any
		Err         string
	}{
		"Invalid filter value": {Input: 1, Expect: 1, Err: "Issue coercing JSON filter"},
		"No matches found":     {Input: "foo", Expect: "foo"},
		"One match found but multiple replacements": { // NEED ALL CAPS DunderVars
			Input: "[__FOO__]", Replacement: "foo|bar", Expect: "[foo]",
		},
		"One replacement but multiple matches": {
			Input: "[__FOO__] && [__BAR__]", Replacement: "fi", Expect: "[fi] && [__BAR__]",
		},
		"All replacements successful": {
			Input: "[__FOO__] && [__BAR__]", Replacement: "foo|bar", Expect: "[foo] && [bar]",
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			jsonObj := map[string]any{"filter_by": testCase.Input}
			os.Setenv("FILTER_REPLACEMENTS", testCase.Replacement)
			err := setFilters(jsonObj)
			finalFilter := jsonObj["filter_by"]
			if finalFilter != testCase.Expect {
				t.Error(test.ErrorMsg("filter", testCase.Expect, finalFilter))
			}
			if !test.IsSameError(err, testCase.Err) {
				t.Error(test.QuotedErrorMsg("error", testCase.Err, err))
			}
		})
	}
}
