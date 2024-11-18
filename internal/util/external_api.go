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

	response, err := client.Post(url, contentType, body) // Sets the "Content-Type" header for you!

	if err != nil {
		return responseData, err
	}

	//INFO: Closing the response body after reading helps the client/server reuse their TCP connection
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return responseData, err
	}

	if err = json.Unmarshal(responseBody, &responseData); err != nil {
		return responseData, err
	}

	return responseData, err
}
