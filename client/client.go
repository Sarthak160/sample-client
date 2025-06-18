package client

import (
	"fmt"
	"io"
	"net/http"
)

type HTTPClient struct{}

func (c *HTTPClient) GetGoogle() (string, error) {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Status: %s, Length: %d", resp.Status, len(body)), nil
}
