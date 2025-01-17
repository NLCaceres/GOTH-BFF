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
		ExpectedResponse string
		ExpectedErr      error
	}{
		"Error within POST itself": { // ASCII Ctrl Char (DEL aka 177) breaks the Server URL
			"/foo" + string([]byte{0x7f}), newHttpMock(403, `{"foo":"bar"}`, nil),
			"", errors.New("parse net/url error"),
		},
		"Error Reading Response": { // Empty response w/ bad StatusCode & Content-Length == 1
			"/foo", newHttpMock(0, ``, map[string]string{"Content-Length": "1"}),
			"", errors.New("unexpected EOF Error"), // Causes this EOF Error
		},
		"Successfully POSTed": {
			"/foo", newHttpMock(202, `{"foo":"bar"}`, nil), `{"foo":"bar"}`, nil,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			server := httptest.NewServer(test.NewTestHandlerFunc(t, testCase.ServerMock))
			defer server.Close()

			serverURL := server.URL + testCase.PostURL
			requestBody := bytes.NewBuffer([]byte(`{"foo":"bar"}`))
			responseBody, err := PostRequest(serverURL, "application/json", requestBody)

			if testCase.ExpectedResponse != string(responseBody) {
				t.Errorf("Response unexpectedly = %v but should be %v", string(responseBody), testCase.ExpectedResponse)
			}
			if (testCase.ExpectedErr == nil && err != nil) || (testCase.ExpectedErr != nil && err == nil) {
				t.Errorf("Error unexpectedly = %v when it should NOT have been", err)
			}
		})
	}
}
func TestPostJSON(t *testing.T) {
	var tests = map[string]struct {
		ServerMock       test.HttpMock
		ExpectedResponse map[string]interface{}
		ExpectedErr      error
	}{
		"Error from internal PostRequest": {
			newHttpMock(0, ``, map[string]string{"Content-Length": "1"}),
			nil, errors.New("parse net/url error"),
		},
		"Error due to Malformed JSON Response": {
			newHttpMock(202, `"foo":"bar"`, nil), // JSON response w/out brackets causes this err
			nil, errors.New("invalid character at top-level of JSON Error"),
		},
		"Successfully POSTed JSON": {
			newHttpMock(202, `{"foo":"bar"}`, nil), map[string]interface{}{"foo": "bar"}, nil,
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			server := httptest.NewServer(test.NewTestHandlerFunc(t, testCase.ServerMock))
			defer server.Close()

			serverURL := server.URL + "/foo"
			requestBody := bytes.NewBuffer([]byte(`{"foo":"bar"}`))
			responseData, err := PostJSON(serverURL, requestBody)
			if testCase.ExpectedResponse == nil && responseData != nil {
				t.Errorf("Response data expected to be nil but was actually filled")
			}
			for key, expectedValue := range testCase.ExpectedResponse {
				if actualValue, ok := responseData[key]; !ok || expectedValue != actualValue {
					t.Errorf("Response map key %v has value of %v instead of %v", key, actualValue, expectedValue)
				}
			}
			if (testCase.ExpectedErr == nil && err != nil) || (testCase.ExpectedErr != nil && err == nil) {
				t.Errorf("Error unexpectedly = %v when it should NOT have been", err)
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
