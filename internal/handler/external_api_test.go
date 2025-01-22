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
	successData := `{"foo":"bar"}`
	tests := map[string]struct {
		Mock               test.HttpMock
		ExpectedStatusCode int
		ExpectedResponse   string
	}{
		"Error from internal Post Request":    {newHttpMock(`"foo":"bar"`), 502, ""},
		"Successfully POSTed to external API": {newHttpMock(successData), 200, successData},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			server := httptest.NewServer(test.NewTestHandlerFunc(t, testCase.Mock))
			defer server.Close()
			os.Setenv("EXTERNAL_API_URL", server.URL)
			// Actual request from client NOT relevant to the test since only grabbing its Path
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
			rec := httptest.NewRecorder()
			c := echo.New().NewContext(req, rec)
			c.SetPath("/foo") //NOTE: BUT must set the Path in Echo's context here!

			ApiPostRequest(c)
			if rec.Code != testCase.ExpectedStatusCode {
				t.Errorf("Response unexpectedly sent %v instead of %v\n", rec.Code, testCase.ExpectedStatusCode)
			}
			if strings.TrimSpace(rec.Body.String()) != testCase.ExpectedResponse {
				t.Errorf("Response unexpectedly sent %v instead of %v\n", rec.Body.String(), testCase.ExpectedResponse)
			}
		})
	}
}

func newHttpMock(data string) test.HttpMock {
	return test.HttpMock{RequestMethod: "POST", ResponseStatus: 200, ResponseData: data}
}
