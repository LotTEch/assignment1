package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// HttpClient er en global HTTP-klient med timeout
var HttpClient = &http.Client{
	Timeout: 10 * time.Second,
}

// DoPostJSON utfører en POST med JSON-body
func DoPostJSON(url string, body interface{}) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	return HttpClient.Do(req)
}
