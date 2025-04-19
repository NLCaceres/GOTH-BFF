package test

import (
	"net/http"
	"testing"
)

type HttpMock struct { // Lowercased Struct fields ARE package private!
	RequestMethod   string
	RequestURL      string
	ResponseStatus  int
	ResponseData    string
	ResponseHeaders map[string]string
}

// TODO: Notable Go Convention fixes to check and make:
// - Drop `New/Get` prefixes from `NewTestHandlerFunc` & `GetProjectFile`
// - Replace `reflect.DeepEquals` with `cmp.Equals` in tests
// - Embrace zero initialiers (string's "", int's 0, bool's false, etc)
// like `new(SomeStruct)` over `&SomeStruct{}` if initial values aren't important
func HttpHandlerFunc(t *testing.T, mock HttpMock) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
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
	}
}
