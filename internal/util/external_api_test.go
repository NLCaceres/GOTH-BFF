package util

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type HttpMock struct { // Lowercased Struct fields ARE package private! Even if it'd be fine here
	RequestMethod   string
	RequestURL      string
	ResponseStatus  int
	ResponseData    string
	ResponseHeaders map[string]string
}

func createHandlerFunc(t *testing.T, mock HttpMock) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != mock.RequestMethod {
			t.Errorf("Unexpected Request method = %s", r.Method)
		}
		// Following could be useful BUT not really interested in checking the URL
		// if r.URL.Path != mock.RequestURL {
		// 	t.Errorf("Unexpected POST Request to URL = %s", r.URL.Path)
		// }
		if mock.ResponseStatus > 0 {
			rw.WriteHeader(mock.ResponseStatus)
		}
		if mock.ResponseData != "" {
			rw.Write([]byte(mock.ResponseData))
		}
		for key, value := range mock.ResponseHeaders {
			rw.Header().Set(key, value)
		}
	})
}

func TestPostRequest(t *testing.T) {
	var tests = map[string]struct {
		PostURL          string
		ServerMock       HttpMock
		ExpectedResponse map[string]interface{}
		ExpectedErr      error
	}{
		"Error within POST itself": { // Using ASCII Ctrl Char (DEL aka 177) to break Server URL
			"/foo" + string([]byte{0x7f}), HttpMock{"POST", "/foo", 403, `{"foo":"bar"}`, map[string]string{}},
			nil, errors.New("parse net/url error"),
		},
		"Error Reading Response": { // An empty Response w/out a StatusCode BUT a Content-Length header of "1"
			"/foo", HttpMock{"POST", "/foo", 0, ``, map[string]string{"Content-Length": "1"}},
			nil, errors.New("unexpected EOF Error"), // Causes this EOF Error
		},
		"Error due to Malformed JSON Response": { // No brackets surrounding the JSON response causes this error
			"/foo", HttpMock{"POST", "/foo", 202, `"foo":"bar"`, map[string]string{}},
			nil, errors.New("invalid character at top-level of JSON Error"),
		},
		"Successfully POSTed": {
			"/foo", HttpMock{"POST", "/foo", 202, `{"foo":"bar"}`, map[string]string{}},
			map[string]interface{}{"foo": "bar"}, nil,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			server := httptest.NewServer(createHandlerFunc(t, test.ServerMock))
			defer server.Close()

			serverURL := server.URL + test.PostURL
			responseData, err := PostRequest(serverURL, "application/json", bytes.NewBuffer([]byte(`{"foo":"bar"}`)))

			if test.ExpectedResponse == nil && responseData != nil {
				t.Errorf("Response data expected to be nil but was actually filled")
			}
			for key, expectedValue := range test.ExpectedResponse {
				if actualValue, ok := responseData[key]; !ok || expectedValue != actualValue {
					t.Errorf("Response map key %v found value of %v instead of %v", key, actualValue, expectedValue)
				}
			}
			if (test.ExpectedErr == nil && err != nil) || (test.ExpectedErr != nil && err == nil) {
				t.Errorf("Error unexpectedly = %v when it shouldn't have been", err)
			}
		})
	}
}
