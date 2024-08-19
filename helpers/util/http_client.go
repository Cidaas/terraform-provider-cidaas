package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HTTPClient struct {
	Token      string
	HTTPMethod string
	URL        string
	Headers    map[string]string
	Host       string
}

type HTTPClientInterface interface {
	MakeRequest(body interface{}) (*http.Response, error)
	SetToken(token string)
	SetMethod(method string)
	SetURL(url string)
	SetHeaders(headers map[string]string)
	GetHost() string
}

func NewHTTPClient(url, method string, token ...string) *HTTPClient {
	var tokenValue string
	if len(token) > 0 {
		tokenValue = token[0]
	}
	return &HTTPClient{
		URL:        url,
		HTTPMethod: method,
		Token:      tokenValue,
	}
}

func (h *HTTPClient) MakeRequest(requestBody interface{}) (*http.Response, error) {
	var reqBodyByte io.Reader
	if requestBody == nil {
		reqBodyByte = nil
	} else {
		bodyByte, err := json.Marshal(requestBody)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal json body, %w", err)
		}
		reqBodyByte = bytes.NewBuffer(bodyByte)
	}

	client := http.DefaultClient
	req, err := http.NewRequest(h.HTTPMethod, h.URL, reqBodyByte)
	if err != nil {
		return nil, fmt.Errorf("failed to create request, %w", err)
	}
	if h.Token != "" {
		req.Header.Add("Authorization", "Bearer "+h.Token)
	}
	for k, v := range h.Headers {
		req.Header.Add(k, v)
	}
	if req.Header.Get("Content-Type") == "" {
		req.Header.Add("Content-Type", "application/json")
	}
	resp, err := client.Do(req)
	if err != nil {
		return resp, fmt.Errorf("request failed, %w", err)
	}
	if (h.HTTPMethod == http.MethodGet || h.HTTPMethod == http.MethodPut) && resp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("status code not 200. found %v, response body: %s", resp.StatusCode, responseToStringConvert(resp))
	}
	if h.HTTPMethod == http.MethodPost && (resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusNoContent) {
		return resp, fmt.Errorf("status code not 20X. found %v, response body: %s", resp.StatusCode, responseToStringConvert(resp))
	}
	if h.HTTPMethod == http.MethodDelete && (resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusAccepted) {
		return resp, fmt.Errorf("status code not 20X. found %v, response body: %s", resp.StatusCode, responseToStringConvert(resp))
	}
	return resp, nil
}

func (h *HTTPClient) SetToken(token string) {
	h.Token = token
}

func (h *HTTPClient) SetMethod(method string) {
	h.HTTPMethod = method
}

func (h *HTTPClient) SetURL(url string) {
	h.URL = url
}

func (h *HTTPClient) SetHeaders(headers map[string]string) {
	h.Headers = headers
}

func (h *HTTPClient) SetHost(host string) {
	h.Host = host
}

func (h *HTTPClient) GetHost() string {
	return h.Host
}
