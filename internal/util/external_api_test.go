package util

import (
	"bytes"
	"errors"
	"github.com/NLCaceres/goth-example/internal/util/test"
	"net/http/httptest"
	"testing"
)

func TestPostRequest(t *testing.T) {
	var tests = map[string]struct {
		PostURL          string
		ServerMock       test.HttpMock
		ExpectedResponse map[string]interface{}
		ExpectedErr      error
	}{
		"Error within POST itself": { // Using ASCII Ctrl Char (DEL aka 177) breaks the Server URL
			"/foo" + string([]byte{0x7f}), newHttpMock(403, `{"foo":"bar"}`, nil),
			nil, errors.New("parse net/url error"),
		},
		"Error Reading Response": { // An empty Response w/out a StatusCode
			"/foo", newHttpMock(0, ``, map[string]string{"Content-Length": "1"}), // BUT a Content-Length header of "1"
			nil, errors.New("unexpected EOF Error"), // causes this EOF Error
		},
		"Error due to Malformed JSON Response": {
			"/foo", newHttpMock(202, `"foo":"bar"`, nil), // No brackets surrounding JSON response
			nil, errors.New("invalid character at top-level of JSON Error"), // Causes this error
		},
		"Successfully POSTed": {
			"/foo", newHttpMock(202, `{"foo":"bar"}`, nil),
			map[string]interface{}{"foo": "bar"}, nil,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			server := httptest.NewServer(test.NewTestHandlerFunc(t, testCase.ServerMock))
			defer server.Close()

			serverURL := server.URL + testCase.PostURL
			responseData, err := PostRequest(serverURL, "application/json", bytes.NewBuffer([]byte(`{"foo":"bar"}`)))

			if testCase.ExpectedResponse == nil && responseData != nil {
				t.Errorf("Response data expected to be nil but was actually filled")
			}
			for key, expectedValue := range testCase.ExpectedResponse {
				if actualValue, ok := responseData[key]; !ok || expectedValue != actualValue {
					t.Errorf("Response map key %v found value of %v instead of %v", key, actualValue, expectedValue)
				}
			}
			if (testCase.ExpectedErr == nil && err != nil) || (testCase.ExpectedErr != nil && err == nil) {
				t.Errorf("Error unexpectedly = %v when it shouldn't have been", err)
			}
		})
	}
}

// Package private helper to condense HttpMock initialization
// Particularly the headers, which when set to `nil` seemingly equal `map[string]string{}`
func newHttpMock(statusCode int, data string, headers map[string]string) test.HttpMock {
	return test.HttpMock{ //NOTE: Structs from other packages REQUIRED field names to protect from breaking changes
		RequestMethod: "POST", RequestURL: "/foo",
		ResponseStatus: statusCode, ResponseData: data, ResponseHeaders: headers,
	}
}
