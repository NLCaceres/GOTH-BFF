package proxy

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func PostRequest(url string, contentType string, body io.Reader) ([]byte, error) {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	client.Timeout = time.Second * 15

	response, err := client.Post(url, contentType, body) // DOES set "Content-Type" header

	if err != nil {
		return nil, RequestError{Cause: err.Error()}
	}

	//INFO: Closing the body after Read helps client/server reuse the TCP connection
	defer response.Body.Close() // Especially if using json.NewDecoder for JSON Streams
	// Instead of Unmarshal for a normal JSON response albeit costing a bit more memory
	responseBody, err := io.ReadAll(response.Body) // by reading the whole body
	if err != nil {
		return nil, RequestError{Cause: err.Error()}
	}

	return responseBody, err
}

// Fires an HTTP POST request delivering the body input, typically a `*bytes.Buffer`,
// to the given URL. The return type represents an HTTP response containing JSON that can
// be unmarshaled into the type input into the generic type parameter.
func PostJSON[T any](url string, body io.Reader) (T, error) {
	var responseData T

	responseBody, err := PostRequest(url, "application/json", body)
	if err != nil {
		return responseData, err
	}

	if err = json.Unmarshal(responseBody, &responseData); err != nil {
		return responseData, RequestError{Cause: err.Error()}
	}

	return responseData, err
}
