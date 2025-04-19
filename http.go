package simple

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var httpClient = &http.Client{}

type HttpClient struct {
	Ctx     context.Context
	Method  string
	Url     string
	Headers map[string]string
	Body    interface{}
	Timeout time.Duration
}

func (c *HttpClient) Request(r interface{}) error {

	c.fill_defaults()

	body, err := c.makeRequest()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, r); err != nil {
		return err
	}

	return nil
}

func (c *HttpClient) fill_defaults() {

	if c.Ctx == nil {
		c.Ctx = context.Background()
	}
	if c.Timeout == 0 {
		c.Timeout = time.Second * 30
	}
	if c.Url == "" {
		c.Url = "http://localhost:8080"
	}
	if c.Headers == nil {
		c.Headers = make(map[string]string)
	}
}

func (c *HttpClient) makeRequest() ([]byte, error) {

	body, err := c.getIOBody()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(c.Ctx, c.Method, c.Url, body)
	if err != nil {
		return nil, err
	}

	c.addHeaders(req)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	code := resp.StatusCode
	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("Request to %s failed with status code %d. Response: %s", c.Url, code, string(respBody))
	}

	return respBody, nil
}

func (c *HttpClient) addHeaders(req *http.Request) {

	if c.Body != nil && c.Headers["Content-Type"] == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	for key, value := range c.Headers {
		req.Header.Add(key, value)
	}
}

func (c *HttpClient) getIOBody() (io.Reader, error) {
	var body io.Reader

	if c.Body != nil {
		jsonBody, err := json.Marshal(c.Body)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(jsonBody)
	}
	return body, nil
}
