package util

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func PostRequest(url string, contentType string, body io.Reader) (map[string]interface{}, error) {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	client.Timeout = time.Second * 15

	var responseData map[string]interface{}

	response, err := client.Post(url, contentType, body) // DOES set "Content-Type" header

	if err != nil {
		return responseData, err
	}

	//INFO: Closing the body after Read helps client/server reuse the TCP connection
	defer response.Body.Close() // Especially if using json.NewDecoder for JSON Streams
	// Instead of Unmarshal for a normal JSON response albeit costing a bit more memory
	responseBody, err := io.ReadAll(response.Body) // by reading the whole body
	if err != nil {
		return responseData, err
	}

	if err = json.Unmarshal(responseBody, &responseData); err != nil {
		return responseData, err
	}

	return responseData, err
}
