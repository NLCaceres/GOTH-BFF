package proxy

import (
	"bytes"
	"github.com/NLCaceres/goth-example/internal/util/test"
	"github.com/google/go-cmp/cmp"
	"net/http/httptest"
	"testing"
)

func TestPostRequest(t *testing.T) {
	tests := map[string]struct {
		PostURL    string
		ServerMock test.HttpMock
		Expect     string
		Err        string
	}{
		"Error within POST itself": { // ASCII Ctrl Char (DEL aka 177) breaks the Server URL
			"/foo" + string([]byte{0x7f}), httpMock(403, `{"foo":"bar"}`, nil), "", "invalid control character",
		},
		"Error Reading Response": { // Empty response w/ bad StatusCode & Content-Length == 1 to trigger error
			"/foo", httpMock(0, ``, map[string]string{"Content-Length": "1"}), "", "unexpected EOF",
		},
		"Successfully POSTed": {
			"/foo", httpMock(202, `{"foo":"bar"}`, nil), `{"foo":"bar"}`, "",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			server := httptest.NewServer(test.HttpHandlerFunc(t, testCase.ServerMock))
			defer server.Close()

			serverURL := server.URL + testCase.PostURL
			requestBody := bytes.NewBuffer([]byte(`{"foo":"bar"}`))
			responseBody, err := PostRequest(serverURL, "application/json", requestBody)

			if testCase.Expect != string(responseBody) {
				t.Errorf("Expected response body = %q but got %q\n", testCase.Expect, string(responseBody))
			}
			if !test.IsSameError(err, testCase.Err) {
				t.Errorf("Expected err = %q but got %q\n", testCase.Err, err)
			}
		})
	}
}
func TestPostJSON(t *testing.T) {
	tests := map[string]struct {
		ServerMock test.HttpMock
		Expect     map[string]any
		Err        string
	}{
		"Error from internal PostRequest": {
			httpMock(0, ``, map[string]string{"Content-Length": "1"}), nil, "unexpected EOF",
		},
		"Error due to Malformed JSON Response": { // Missing brackets in JSON response to trigger err
			httpMock(202, `"foo":"bar"`, nil), nil, "invalid character",
		},
		"Successfully POSTed JSON": {
			httpMock(202, `{"foo":"bar"}`, nil), map[string]any{"foo": "bar"}, "",
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			server := httptest.NewServer(test.HttpHandlerFunc(t, testCase.ServerMock))
			defer server.Close()

			serverURL := server.URL + "/foo"
			requestBody := bytes.NewBuffer([]byte(`{"foo":"bar"}`))
			responseData, err := PostJSON(serverURL, requestBody)
			if !cmp.Equal(testCase.Expect, responseData) {
				t.Error(test.ErrorMsg("response data", testCase.Expect, responseData))
			}
			if !test.IsSameError(err, testCase.Err) {
				t.Errorf("Expected err = %q but got %q\n", testCase.Err, err)
			}
		})
	}
}

// Package private helper to condense HttpMock initialization
// Particularly the headers, which when set to `nil` seemingly equal `map[string]string{}`
func httpMock(statusCode int, data string, headers map[string]string) test.HttpMock {
	return test.HttpMock{ //NOTE: Other pkg Structs MUST use FieldNames to protect against changes
		RequestMethod: "POST", RequestURL: "/foo",
		ResponseStatus: statusCode, ResponseData: data, ResponseHeaders: headers,
	}
}
