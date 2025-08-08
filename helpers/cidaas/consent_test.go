// helpers/cidaas/consent_test.go
package cidaas

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewConsent(t *testing.T) {
	config := NewTestClientConfig("http://test.com")

	consent := NewConsent(config)

	if consent == nil {
		t.Fatal("Expected consent instance, got nil")
	}

	if consent.BaseURL != config.BaseURL {
		t.Errorf("Expected BaseURL %s, got %s", config.BaseURL, consent.BaseURL)
	}

	if consent.AccessToken != config.AccessToken {
		t.Errorf("Expected AccessToken %s, got %s", config.AccessToken, consent.AccessToken)
	}
}

func TestConsent_Upsert_Success(t *testing.T) {
	expectedConsent := ConsentModel{
		ID:             "consent-123",
		ConsentGroupID: "group-456",
		ConsentName:    "Data Processing Consent",
		Enabled:        true,
		CreatedTime:    "2024-01-15T10:30:00Z",
		UpdatedTime:    "2024-01-15T11:45:00Z",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, "consent-management-srv/v2/consent/instance") {
			t.Errorf("Expected consent-management-srv/v2/consent/instance endpoint, got %s", r.URL.Path)
		}

		body, _ := io.ReadAll(r.Body)
		var receivedConsent ConsentModel
		json.Unmarshal(body, &receivedConsent)

		if receivedConsent.ConsentName != expectedConsent.ConsentName {
			t.Errorf("Expected ConsentName %s, got %s", expectedConsent.ConsentName, receivedConsent.ConsentName)
		}

		if receivedConsent.Enabled != expectedConsent.Enabled {
			t.Errorf("Expected Enabled %t, got %t", expectedConsent.Enabled, receivedConsent.Enabled)
		}

		response := ConsentResponse{
			Success: true,
			Status:  201,
			Data:    expectedConsent,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consent := NewConsent(config)

	result, err := consent.Upsert(context.Background(), expectedConsent)

	if err != nil {
		t.Fatalf("Upsert failed: %v", err)
	}

	if !result.Success {
		t.Error("Expected success to be true")
	}

	if result.Data.ConsentName != expectedConsent.ConsentName {
		t.Errorf("Expected ConsentName %s, got %s", expectedConsent.ConsentName, result.Data.ConsentName)
	}

	if result.Data.Enabled != expectedConsent.Enabled {
		t.Errorf("Expected Enabled %t, got %t", expectedConsent.Enabled, result.Data.Enabled)
	}
}

func TestConsent_Upsert_Error(t *testing.T) {
	server := NewMockServer(http.StatusBadRequest, `{"error": "invalid consent"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consent := NewConsent(config)

	consentModel := ConsentModel{ConsentName: "invalid"}
	_, err := consent.Upsert(context.Background(), consentModel)

	if err == nil {
		t.Error("Expected error for bad request, got nil")
	}
}

func TestConsent_GetConsentInstances_Success(t *testing.T) {
	expectedConsents := []ConsentModel{
		{
			ID:             "consent-1",
			ConsentGroupID: "group-123",
			ConsentName:    "Marketing Consent",
			Enabled:        true,
		},
		{
			ID:             "consent-2",
			ConsentGroupID: "group-123",
			ConsentName:    "Analytics Consent",
			Enabled:        false,
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, "consent-management-srv/v2/consent/instance/group-123") {
			t.Errorf("Expected consent-management-srv/v2/consent/instance/group-123 endpoint, got %s", r.URL.Path)
		}

		response := ConsentInstanceResponse{
			Success: true,
			Status:  200,
			Data:    expectedConsents,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consent := NewConsent(config)

	result, err := consent.GetConsentInstances(context.Background(), "group-123")

	if err != nil {
		t.Fatalf("GetConsentInstances failed: %v", err)
	}

	if !result.Success {
		t.Error("Expected success to be true")
	}

	if len(result.Data) != 2 {
		t.Errorf("Expected 2 consent instances, got %d", len(result.Data))
	}

	if result.Data[0].ConsentName != "Marketing Consent" {
		t.Errorf("Expected first consent name 'Marketing Consent', got %s", result.Data[0].ConsentName)
	}
}

func TestConsent_GetConsentInstances_NoContent(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consent := NewConsent(config)

	result, err := consent.GetConsentInstances(context.Background(), "empty-group")

	if err != nil {
		t.Fatalf("GetConsentInstances with no content failed: %v", err)
	}

	if result.Success {
		t.Error("Expected success to be false for no content")
	}

	if result.Status != http.StatusNoContent {
		t.Errorf("Expected status %d, got %d", http.StatusNoContent, result.Status)
	}
}

func TestConsent_GetConsentInstances_Error(t *testing.T) {
	server := NewMockServer(http.StatusNotFound, `{"error": "consent group not found"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consent := NewConsent(config)

	_, err := consent.GetConsentInstances(context.Background(), "nonexistent")

	if err == nil {
		t.Error("Expected error for not found consent group, got nil")
	}
}

func TestConsent_Delete_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, "consent-management-srv/v2/consent/instance/consent-123") {
			t.Errorf("Expected consent-management-srv/v2/consent/instance/consent-123 endpoint, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consent := NewConsent(config)

	err := consent.Delete(context.Background(), "consent-123")

	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
}

func TestConsent_Delete_Error(t *testing.T) {
	server := NewMockServer(http.StatusNotFound, `{"error": "consent not found"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consent := NewConsent(config)

	err := consent.Delete(context.Background(), "nonexistent")

	if err == nil {
		t.Error("Expected error for not found consent, got nil")
	}
}

func TestConsent_GetAll_Success(t *testing.T) {
	expectedConsents := []ConsentModel{
		{
			ID:             "consent-1",
			ConsentGroupID: "group-1",
			ConsentName:    "Privacy Consent",
			Enabled:        true,
			CreatedTime:    "2024-01-10T09:00:00Z",
		},
		{
			ID:             "consent-2",
			ConsentGroupID: "group-2",
			ConsentName:    "Cookie Consent",
			Enabled:        false,
			CreatedTime:    "2024-01-11T10:00:00Z",
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, "consent-management-srv/v2/consent/instance/all/list") {
			t.Errorf("Expected consent-management-srv/v2/consent/instance/all/list endpoint, got %s", r.URL.Path)
		}

		response := ConsentInstanceResponse{
			Success: true,
			Status:  200,
			Data:    expectedConsents,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consent := NewConsent(config)

	result, err := consent.GetAll(context.Background())

	if err != nil {
		t.Fatalf("GetAll failed: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("Expected 2 consents, got %d", len(result))
	}

	if result[0].ConsentName != "Privacy Consent" {
		t.Errorf("Expected first consent name 'Privacy Consent', got %s", result[0].ConsentName)
	}

	if result[1].ConsentName != "Cookie Consent" {
		t.Errorf("Expected second consent name 'Cookie Consent', got %s", result[1].ConsentName)
	}
}

func TestConsent_GetAll_EmptyResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := ConsentInstanceResponse{
			Success: true,
			Status:  200,
			Data:    []ConsentModel{},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consent := NewConsent(config)

	result, err := consent.GetAll(context.Background())

	if err != nil {
		t.Fatalf("GetAll failed: %v", err)
	}

	if len(result) != 0 {
		t.Errorf("Expected 0 consents, got %d", len(result))
	}
}

func TestConsent_GetAll_Error(t *testing.T) {
	server := NewMockServer(http.StatusInternalServerError, `{"error": "server error"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consent := NewConsent(config)

	_, err := consent.GetAll(context.Background())

	if err == nil {
		t.Error("Expected error for server error, got nil")
	}
}

func TestConsent_Upsert_DisabledConsent(t *testing.T) {
	disabledConsent := ConsentModel{
		ConsentGroupID: "group-789",
		ConsentName:    "Disabled Consent",
		Enabled:        false,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var receivedConsent ConsentModel
		json.Unmarshal(body, &receivedConsent)

		if receivedConsent.Enabled {
			t.Error("Expected Enabled to be false")
		}

		response := ConsentResponse{
			Success: true,
			Status:  201,
			Data:    disabledConsent,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consent := NewConsent(config)

	result, err := consent.Upsert(context.Background(), disabledConsent)

	if err != nil {
		t.Fatalf("Upsert disabled consent failed: %v", err)
	}

	if result.Data.Enabled {
		t.Error("Expected Enabled to be false in response")
	}
}
