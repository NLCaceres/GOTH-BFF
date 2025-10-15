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

func TestNewQuery(t *testing.T) {
	badData := `"foo":"bar"`
	successData := "{" + badData + "}"
	tests := map[string]struct {
		Mock               test.HttpMock
		QueryFile          string
		Filters            string
		ExpectedStatusCode int
		ExpectedResponse   string
	}{
		"Error reading unknown JSON":    {httpMock(badData), "./bad.json", "", 500, ""},
		"Error due to bad filter value": {httpMock(badData), "./bad_typing.json", "", 500, ""},
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
			NewQuery(c)
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
