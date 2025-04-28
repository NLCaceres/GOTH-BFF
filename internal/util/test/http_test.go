package test

import (
	"bytes"
	"net/http"
	"testing"
)

func TestHttpHandlerFunc(t *testing.T) {
	tests := map[string]struct {
		Input       HttpMock
		ExpectCount int // Tracks function calling like a test spy
	}{
		"Calls WriteHeader, Write & Header ONCE": {
			httpMock(200, `{"foo":"bar"}`, map[string]string{"Content-Length": "1"}), 1,
		},
		"DOESN'T Call WriteHeader, Write, or Header": {
			httpMock(0, ``, nil), 0,
		}, // NOTE: Adding a test case for when Request Methods ("POST, "GET", etc) don't match ALWAYS fails
	} // due to underlying `t.Error()` AND the fail CAN'T be prevented, so maybe rewrite Handler later?

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			handler := HttpHandlerFunc(t, testCase.Input)
			mockRequest, _ := http.NewRequest(testCase.Input.RequestMethod, testCase.Input.RequestURL, bytes.NewBuffer([]byte(``)))
			// Setting the Map keys ensures it's checked when the funcs under test are NOT called
			mockResponseWriter := mockResponseWriter{t, map[string]int{"WriteHeader": 0, "Write": 0, "Header": 0}}

			handler(mockResponseWriter, mockRequest)

			for funcName, count := range mockResponseWriter.CallCounter {
				if count != testCase.ExpectCount {
					t.Errorf("Expected ResponseWriter %v() called %d times but got %d", funcName, testCase.ExpectCount, count)
				}
			}
		})
	}
}

// NOTE: Package private Mocks with sensible defaults
type mockResponseWriter struct {
	TestHelper  *testing.T
	CallCounter map[string]int
}

func (rw mockResponseWriter) Header() http.Header {
	rw.CallCounter["Header"]++
	return map[string][]string{}
}
func (rw mockResponseWriter) Write([]byte) (int, error) {
	rw.CallCounter["Write"]++
	return 1, nil
}
func (rw mockResponseWriter) WriteHeader(statusCode int) {
	rw.CallCounter["WriteHeader"]++
}

func httpMock(statusCode int, responseData string, responseHeaders map[string]string) HttpMock {
	return HttpMock{"GET", "/fizz", statusCode, responseData, responseHeaders}
}
