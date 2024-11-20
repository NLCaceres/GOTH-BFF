package test

import (
	"bytes"
	"net/http"
	"testing"
)

func TestNewTestHandlerFunc(t *testing.T) {
	var tests = map[string]struct {
		input         HttpMock
		expectedCount int
	}{
		"Calls WriteHeader, Write & Header ONCE": {
			newHttpMock(200, `{"foo":"bar"}`, map[string]string{"Content-Length": "1"}), 1,
		},
		"DOESN'T Call WriteHeader, Write, or Header": {
			newHttpMock(0, ``, nil), 0,
		}, // NOTE: Adding a test case for when Request Methods ("POST, "GET", etc) don't match ALWAYS fails
	} // due to underlying `t.Error()` AND the fail CAN'T be prevented, so maybe rewrite Handler later?

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			handler := NewTestHandlerFunc(t, test.input)
			mockRequest, _ := http.NewRequest(test.input.RequestMethod, test.input.RequestURL, bytes.NewBuffer([]byte(``)))
			// Setting the Map keys ensures it's checked when the funcs under test are NOT called
			mockResponseWriter := mockResponseWriter{t, map[string]int{"WriteHeader": 0, "Write": 0, "Header": 0}}

			handler(mockResponseWriter, mockRequest)

			for funcName, count := range mockResponseWriter.CallCounter {
				if count != test.expectedCount {
					t.Errorf("ResponseWriter %v() unexpectedly called %d times BUT expected %d times", funcName, count, test.expectedCount)
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

func newHttpMock(statusCode int, responseData string, responseHeaders map[string]string) HttpMock {
	return HttpMock{"GET", "/fizz", statusCode, responseData, responseHeaders}
}
