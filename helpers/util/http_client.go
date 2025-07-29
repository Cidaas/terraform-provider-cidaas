package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// HTTPClient provides a configurable HTTP client with authentication and error handling.
// It supports common HTTP methods and automatic JSON marshaling.
type HTTPClient struct {
	Token      string
	HTTPMethod string
	URL        string
	Headers    map[string]string
}

type HTTPClientInterface interface {
	MakeRequest(body interface{}) (*http.Response, error)
}

// NewHTTPClient creates a new HTTP client with the specified URL and method.
// Optional token parameter enables Bearer authentication.
func NewHTTPClient(url, method string, token ...string) (*HTTPClient, error) {
	if url == "" {
		return nil, fmt.Errorf("URL cannot be empty")
	}
	if method == "" {
		return nil, fmt.Errorf("HTTP method cannot be empty")
	}

	var tokenValue string
	if len(token) > 0 {
		tokenValue = token[0]
	}
	return &HTTPClient{
		URL:        url,
		HTTPMethod: method,
		Token:      tokenValue,
		Headers:    make(map[string]string),
	}, nil
}

// MakeRequest executes an HTTP request with the configured method, URL, and headers.
// It automatically marshals the request body to JSON and handles authentication if a token is set.
func (h *HTTPClient) MakeRequest(ctx context.Context, requestBody interface{}) (*http.Response, error) {
	var reqBodyByte io.Reader
	if requestBody == nil {
		reqBodyByte = nil
	} else {
		bodyByte, err := json.Marshal(requestBody)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal JSON body, %w", err)
		}
		reqBodyByte = bytes.NewBuffer(bodyByte)
	}

	client := http.DefaultClient
	req, err := http.NewRequestWithContext(ctx, h.HTTPMethod, h.URL, reqBodyByte)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request, %w", err)
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

	var expectedCodes []int
	switch h.HTTPMethod {
	case http.MethodGet, http.MethodPut:
		expectedCodes = []int{http.StatusOK}
	case http.MethodPost:
		expectedCodes = []int{http.StatusOK, http.StatusCreated, http.StatusNoContent}
	case http.MethodDelete:
		expectedCodes = []int{http.StatusOK, http.StatusCreated, http.StatusNoContent, http.StatusAccepted}
	default:
		expectedCodes = []int{http.StatusOK} // Default fallback
	}

	if err := h.handleErrorResponse(resp, expectedCodes); err != nil {
		return resp, err
	}
	return resp, nil
}

// handleErrorResponse validates the HTTP response status code against expected codes.
// Returns an error if the status code is not in the expected list, including response body for debugging.
func (h *HTTPClient) handleErrorResponse(resp *http.Response, expectedCodes []int) error {
	for _, code := range expectedCodes {
		if resp.StatusCode == code {
			return nil
		}
	}

	bodyStr, bodyErr := ResponseToString(resp)
	if bodyErr != nil {
		return fmt.Errorf("unexpected status code %d, failed to read response body: %w", resp.StatusCode, bodyErr)
	}
	return fmt.Errorf("unexpected status code %d, response body: %s", resp.StatusCode, bodyStr)
}
