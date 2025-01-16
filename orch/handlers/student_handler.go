package handlers

import (
	"fmt"
	"io"
	"net/http"
)

const coreBaseURL = "http://localhost:3002" // Core service URL

func sendRequestToCoreValidate(method, endpoint string, body io.Reader) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", coreBaseURL, endpoint)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}

func readResponseBodyValidate(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
