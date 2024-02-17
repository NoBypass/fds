package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	url   string
	token string
}

// NewFDSClient creates a new client for the FDS API.
// For most cases, it is recommended to use the
// SetToken method soon after to set the token.
func NewFDSClient(url string) *Client {
	return &Client{
		url: url,
	}
}

// SetToken sets the token for the client. This token
// for is used for most endpoints.
func (c *Client) SetToken(token string) {
	c.token = token
}

func (c *Client) newJsonRequest(method, path string, body interface{}) (*http.Request, error) {
	var reader *bytes.Reader = nil
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewReader(data)
	}
	req, err := http.NewRequest(method, c.url+path, reader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)
	return req, nil
}

func do[T any](req *http.Request) (*T, error) {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, err
	}

	v := new(T)
	err = json.NewDecoder(res.Body).Decode(v)
	return v, err
}
