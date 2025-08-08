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

func TestNewConsentGroup(t *testing.T) {
	config := NewTestClientConfig("http://test.com")

	consentGroup := NewConsentGroup(config)

	if consentGroup == nil {
		t.Fatal("Expected consent group instance, got nil")
	}

	if consentGroup.BaseURL != config.BaseURL {
		t.Errorf("Expected BaseURL %s, got %s", config.BaseURL, consentGroup.BaseURL)
	}

	if consentGroup.AccessToken != config.AccessToken {
		t.Errorf("Expected AccessToken %s, got %s", config.AccessToken, consentGroup.AccessToken)
	}
}

func TestConsentGroup_Upsert_Success(t *testing.T) {
	expectedConsentGroup := ConsentGroupConfig{
		ID:          "cg-123",
		GroupName:   "Marketing Consent Group",
		Description: "Group for marketing related consents",
		CreatedTime: "2024-01-15T10:30:00Z",
		UpdatedTime: "2024-01-15T11:45:00Z",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, "consent-management-srv/v2/groups") {
			t.Errorf("Expected consent-management-srv/v2/groups endpoint, got %s", r.URL.Path)
		}

		body, _ := io.ReadAll(r.Body)
		var receivedConsentGroup ConsentGroupConfig
		json.Unmarshal(body, &receivedConsentGroup)

		if receivedConsentGroup.GroupName != expectedConsentGroup.GroupName {
			t.Errorf("Expected GroupName %s, got %s", expectedConsentGroup.GroupName, receivedConsentGroup.GroupName)
		}

		if receivedConsentGroup.Description != expectedConsentGroup.Description {
			t.Errorf("Expected Description %s, got %s", expectedConsentGroup.Description, receivedConsentGroup.Description)
		}

		response := ConsentGroupResponse{
			Success: true,
			Status:  201,
			Data:    expectedConsentGroup,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentGroup := NewConsentGroup(config)

	result, err := consentGroup.Upsert(context.Background(), expectedConsentGroup)

	if err != nil {
		t.Fatalf("Upsert failed: %v", err)
	}

	if !result.Success {
		t.Error("Expected success to be true")
	}

	if result.Data.GroupName != expectedConsentGroup.GroupName {
		t.Errorf("Expected GroupName %s, got %s", expectedConsentGroup.GroupName, result.Data.GroupName)
	}

	if result.Data.Description != expectedConsentGroup.Description {
		t.Errorf("Expected Description %s, got %s", expectedConsentGroup.Description, result.Data.Description)
	}
}

func TestConsentGroup_Upsert_WithEmptyDescription(t *testing.T) {
	consentGroupWithEmptyDesc := ConsentGroupConfig{
		GroupName:   "Empty Description Group",
		Description: "", // Testing the special case where description can be empty string
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var receivedConsentGroup ConsentGroupConfig
		json.Unmarshal(body, &receivedConsentGroup)

		// Verify that empty description is preserved (not omitted)
		if receivedConsentGroup.Description != "" {
			t.Errorf("Expected empty Description, got %s", receivedConsentGroup.Description)
		}

		response := ConsentGroupResponse{
			Success: true,
			Status:  201,
			Data:    consentGroupWithEmptyDesc,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentGroup := NewConsentGroup(config)

	result, err := consentGroup.Upsert(context.Background(), consentGroupWithEmptyDesc)

	if err != nil {
		t.Fatalf("Upsert with empty description failed: %v", err)
	}

	if result.Data.Description != "" {
		t.Errorf("Expected empty Description in response, got %s", result.Data.Description)
	}
}

func TestConsentGroup_Upsert_Error(t *testing.T) {
	server := NewMockServer(http.StatusBadRequest, `{"error": "invalid consent group"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentGroup := NewConsentGroup(config)

	consentGroupConfig := ConsentGroupConfig{GroupName: "invalid"}
	_, err := consentGroup.Upsert(context.Background(), consentGroupConfig)

	if err == nil {
		t.Error("Expected error for bad request, got nil")
	}
}

func TestConsentGroup_Get_Success(t *testing.T) {
	expectedConsentGroup := ConsentGroupConfig{
		ID:          "cg-456",
		GroupName:   "Analytics Consent Group",
		Description: "Group for analytics related consents",
		CreatedTime: "2024-01-10T09:00:00Z",
		UpdatedTime: "2024-01-10T09:30:00Z",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, "consent-management-srv/v2/groups/cg-456") {
			t.Errorf("Expected consent-management-srv/v2/groups/cg-456 endpoint, got %s", r.URL.Path)
		}

		response := ConsentGroupResponse{
			Success: true,
			Status:  200,
			Data:    expectedConsentGroup,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentGroup := NewConsentGroup(config)

	result, err := consentGroup.Get(context.Background(), "cg-456")

	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}

	if !result.Success {
		t.Error("Expected success to be true")
	}

	if result.Data.ID != expectedConsentGroup.ID {
		t.Errorf("Expected ID %s, got %s", expectedConsentGroup.ID, result.Data.ID)
	}

	if result.Data.GroupName != expectedConsentGroup.GroupName {
		t.Errorf("Expected GroupName %s, got %s", expectedConsentGroup.GroupName, result.Data.GroupName)
	}
}

func TestConsentGroup_Get_Error(t *testing.T) {
	server := NewMockServer(http.StatusNotFound, `{"error": "consent group not found"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentGroup := NewConsentGroup(config)

	_, err := consentGroup.Get(context.Background(), "nonexistent")

	if err == nil {
		t.Error("Expected error for not found consent group, got nil")
	}
}

func TestConsentGroup_Delete_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, "consent-management-srv/v2/groups/cg-789") {
			t.Errorf("Expected consent-management-srv/v2/groups/cg-789 endpoint, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentGroup := NewConsentGroup(config)

	err := consentGroup.Delete(context.Background(), "cg-789")

	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
}

func TestConsentGroup_Delete_Error(t *testing.T) {
	server := NewMockServer(http.StatusNotFound, `{"error": "consent group not found"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentGroup := NewConsentGroup(config)

	err := consentGroup.Delete(context.Background(), "nonexistent")

	if err == nil {
		t.Error("Expected error for not found consent group, got nil")
	}
}

func TestConsentGroup_Upsert_MinimalData(t *testing.T) {
	minimalConsentGroup := ConsentGroupConfig{
		GroupName:   "Minimal Group",
		Description: "",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var receivedConsentGroup ConsentGroupConfig
		json.Unmarshal(body, &receivedConsentGroup)

		if receivedConsentGroup.GroupName != minimalConsentGroup.GroupName {
			t.Errorf("Expected GroupName %s, got %s", minimalConsentGroup.GroupName, receivedConsentGroup.GroupName)
		}

		// Server might add timestamps and ID
		responseData := minimalConsentGroup
		responseData.ID = "generated-id-123"
		responseData.CreatedTime = "2024-01-15T10:30:00Z"
		responseData.UpdatedTime = "2024-01-15T10:30:00Z"

		response := ConsentGroupResponse{
			Success: true,
			Status:  201,
			Data:    responseData,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentGroup := NewConsentGroup(config)

	result, err := consentGroup.Upsert(context.Background(), minimalConsentGroup)

	if err != nil {
		t.Fatalf("Upsert with minimal data failed: %v", err)
	}

	if !result.Success {
		t.Error("Expected success to be true")
	}

	if result.Data.GroupName != minimalConsentGroup.GroupName {
		t.Errorf("Expected GroupName %s, got %s", minimalConsentGroup.GroupName, result.Data.GroupName)
	}

	// Verify server-generated fields
	if result.Data.ID == "" {
		t.Error("Expected server to generate an ID")
	}

	if result.Data.CreatedTime == "" {
		t.Error("Expected server to set CreatedTime")
	}
}

func TestConsentGroup_Get_WithSpecialCharacters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify URL path contains the special character ID
		if !strings.Contains(r.URL.Path, "consent-group_123") {
			t.Errorf("Expected 'consent-group_123' in URL path, got %s", r.URL.Path)
		}

		response := ConsentGroupResponse{
			Success: true,
			Status:  200,
			Data: ConsentGroupConfig{
				ID:        "consent-group_123",
				GroupName: "Special Characters Group",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentGroup := NewConsentGroup(config)

	_, err := consentGroup.Get(context.Background(), "consent-group_123")

	if err != nil {
		t.Fatalf("Get with special characters failed: %v", err)
	}
}

func TestConsentGroup_Delete_WithSpecialCharacters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify URL path contains the special character ID
		if !strings.Contains(r.URL.Path, "consent-group_456") {
			t.Errorf("Expected 'consent-group_456' in URL path, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentGroup := NewConsentGroup(config)

	err := consentGroup.Delete(context.Background(), "consent-group_456")

	if err != nil {
		t.Fatalf("Delete with special characters failed: %v", err)
	}
}
