package proxy

import (
	"bytes"
	"errors"
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
		Err        error
	}{
		"Error within POST itself": { // ASCII Ctrl Char (DEL aka 177) breaks the Server URL
			"/foo" + string([]byte{0x7f}), httpMock(403, `{"foo":"bar"}`, nil),
			"", errors.New("parse net/url error"),
		},
		"Error Reading Response": { // Empty response w/ bad StatusCode & Content-Length == 1
			"/foo", httpMock(0, ``, map[string]string{"Content-Length": "1"}),
			"", errors.New("unexpected EOF Error"), // Causes this EOF Error
		},
		"Successfully POSTed": {
			"/foo", httpMock(202, `{"foo":"bar"}`, nil), `{"foo":"bar"}`, nil,
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
				t.Errorf("Response unexpectedly = %v but should be %v", string(responseBody), testCase.Expect)
			}
			if test.OnlyOneIsNil(testCase.Err, err) {
				t.Errorf("Error unexpectedly = %v when it should NOT have been", err)
			}
		})
	}
}
func TestPostJSON(t *testing.T) {
	tests := map[string]struct {
		ServerMock test.HttpMock
		Expect     map[string]interface{}
		Err        error
	}{
		"Error from internal PostRequest": {
			httpMock(0, ``, map[string]string{"Content-Length": "1"}),
			nil, errors.New("unexpected EOF Error"),
		},
		"Error due to Malformed JSON Response": {
			httpMock(202, `"foo":"bar"`, nil), // JSON response w/out brackets causes this err
			nil, errors.New("invalid character at top-level of JSON Error"),
		},
		"Successfully POSTed JSON": {
			httpMock(202, `{"foo":"bar"}`, nil), map[string]interface{}{"foo": "bar"}, nil,
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
				t.Errorf("Expected response of %v but got %v", testCase.Expect, responseData)
			}
			if test.OnlyOneIsNil(testCase.Err, err) {
				t.Errorf("Error unexpectedly = %v when it should NOT have been", err)
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
