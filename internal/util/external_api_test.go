package util

import (
	"bytes"
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

func TestPostRequestError(t *testing.T) {
	server := httptest.NewServer(createHandlerFunc(t, HttpMock{"POST", "/foo", 403, `{"foo":"bar"}`, map[string]string{}}))
	defer server.Close()

	// Can break the URL with an ASCII Ctrl Char (0-31 or 177). 0x7f is the 177 DEL Char
	badChar := string([]byte{0x7f}) // Directly converting via string(0x7f) issues a warning
	serverURL := server.URL + "/foo" + badChar
	responseData, err := PostRequest(serverURL, "application/json", bytes.NewBuffer([]byte(`{"foo":"bar"}`)))
	// Still should get a map back, but it'll be empty (and therefore equal to nil)
	if responseData != nil && len(responseData) != 0 {
		t.Errorf("Response map is unexpectedly filled with length of %v", len(responseData))
	}
	if err == nil {
		t.Error("Error was nil, but should have been a parsing error from the `net/url` package")
	}
}

func TestResponseReadError(t *testing.T) {
	// Adding the "Content-Length" Header w/out setting any response data or status causes an EOF error
	server := httptest.NewServer(createHandlerFunc(t, HttpMock{
		"POST", "/foo", 0, ``, map[string]string{"Content-Length": "1"},
	}))
	defer server.Close()

	serverURL := server.URL + "/foo"
	responseData, err := PostRequest(serverURL, "application/json", bytes.NewBuffer([]byte(`{"foo":"bar"}`)))

	if responseData != nil && len(responseData) != 0 {
		t.Errorf("Response map is unexpectedly filled with length of %v", len(responseData))
	}
	if err == nil {
		t.Errorf("Error was nil, but should have been an EOF error")
	}
}

func TestJsonUnmarshalError(t *testing.T) {
	// A badly formatted JSON string, like one without surrounding brackets, causes a JSON error
	server := httptest.NewServer(createHandlerFunc(t, HttpMock{
		"POST", "/foo", 202, `"foo":"bar"`, map[string]string{},
	}))
	defer server.Close()

	serverURL := server.URL + "/foo"
	responseData, err := PostRequest(serverURL, "application/json", bytes.NewBuffer([]byte(`{"foo":"bar"}`)))

	if responseData != nil && len(responseData) != 0 {
		t.Errorf("Response map is unexpectedly filled with length of %v", len(responseData))
	}
	if err == nil {
		t.Errorf("Error was nil, but should have been a JSON error")
	}
}

func TestPostSuccess(t *testing.T) {
	server := httptest.NewServer(createHandlerFunc(t, HttpMock{
		"POST", "/foo", 202, `{"foo":"bar"}`, map[string]string{},
	}))
	defer server.Close()

	serverURL := server.URL + "/foo"
	responseData, err := PostRequest(serverURL, "application/json", bytes.NewBuffer([]byte(`{"foo":"bar"}`)))

	if err != nil {
		t.Errorf("Error expected to be nil but got %v", err)
	}
	if value, ok := responseData["foo"]; responseData == nil || len(responseData) == 0 || !ok || value != "bar" {
		t.Errorf("Response map expected to contain key-val pairs but got %v", responseData)
	}
}
