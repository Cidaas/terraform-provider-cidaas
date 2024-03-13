package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type HttpClient struct {
	Token string
}

type HttpClientInterface interface {
	Get(url string) (*http.Response, error)
	Post(url string, requestBody *strings.Reader) (*http.Response, error)
	Put(url string, requestBody *strings.Reader) (*http.Response, error)
	Delete(url string) (*http.Response, error)
}

func (h *HttpClient) Get(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request, %v", err)
	}
	req.Header.Add("Authorization", "Bearer "+h.Token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed, %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code not 200. found %v, response body: %v", resp.StatusCode, responseToStringConvert(resp))
	}
	return resp, nil
}

func (h *HttpClient) Post(url string, body interface{}) (*http.Response, error) {
	reqBodyByte, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json body, %v", err)
	}
	client := http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBodyByte))
	if err != nil {
		return nil, fmt.Errorf("failed to create request, %v", err)
	}
	if h.Token != "" {
		req.Header.Add("Authorization", "Bearer "+h.Token)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed, %v", err)
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("status code not 20X. found %v, response body: %v", resp.StatusCode, responseToStringConvert(resp))
	}
	return resp, nil
}

func (h *HttpClient) Put(url string, body interface{}) (*http.Response, error) {
	reqBodyByte, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json body, %v", err)
	}
	client := http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(reqBodyByte))
	if err != nil {
		return nil, fmt.Errorf("failed to create request, %v", err)
	}
	req.Header.Add("Authorization", "Bearer "+h.Token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("request failed, %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code not 20X. found %v, response body: %v", resp.StatusCode, responseToStringConvert(resp))
	}
	return resp, nil
}

func (h *HttpClient) Delete(url string) (*http.Response, error) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request, %v", err)
	}
	req.Header.Add("Authorization", "Bearer "+h.Token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed, %v", err)
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf("status code not 20X. found %v, response body: %v", resp.StatusCode, responseToStringConvert(resp))
	}
	return resp, nil
}
