// helpers/cidaas/consent_version_test.go
package cidaas

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewConsentVersion(t *testing.T) {
	config := NewTestClientConfig("http://test.com")

	consentVersion := NewConsentVersion(config)

	if consentVersion == nil {
		t.Fatal("Expected consent version instance, got nil")
	}

	if consentVersion.BaseURL != config.BaseURL {
		t.Errorf("Expected BaseURL %s, got %s", config.BaseURL, consentVersion.BaseURL)
	}

	if consentVersion.AccessToken != config.AccessToken {
		t.Errorf("Expected AccessToken %s, got %s", config.AccessToken, consentVersion.AccessToken)
	}
}

func TestConsentVersion_Upsert_Success(t *testing.T) {
	expectedConsentVersion := ConsentVersionModel{
		ID:             "cv-123",
		Version:        1.5,
		ConsentID:      "consent-456",
		ConsentType:    "PRIVACY",
		Scopes:         []string{"profile", "email", "phone"},
		RequiredFields: []string{"email", "given_name", "family_name"},
		ConsentLocale: ConsentLocale{
			Locale:  "en-US",
			Content: "This is the privacy consent content in English.",
			URL:     "https://example.com/privacy-en",
		},
		CreatedAt: "2024-01-15T10:30:00Z",
		UpdatedAt: "2024-01-15T11:45:00Z",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, "consent-management-srv/v2/consent/versions") {
			t.Errorf("Expected consent-management-srv/v2/consent/versions endpoint, got %s", r.URL.Path)
		}

		// Verify Authorization header
		authHeader := r.Header.Get("Authorization")
		if !strings.Contains(authHeader, "test-token") {
			t.Errorf("Expected Authorization header with token, got %s", authHeader)
		}

		body, _ := io.ReadAll(r.Body)
		var receivedConsentVersion ConsentVersionModel
		json.Unmarshal(body, &receivedConsentVersion)

		if receivedConsentVersion.ConsentID != expectedConsentVersion.ConsentID {
			t.Errorf("Expected ConsentID %s, got %s", expectedConsentVersion.ConsentID, receivedConsentVersion.ConsentID)
		}

		if receivedConsentVersion.Version != expectedConsentVersion.Version {
			t.Errorf("Expected Version %f, got %f", expectedConsentVersion.Version, receivedConsentVersion.Version)
		}

		if len(receivedConsentVersion.Scopes) != 3 {
			t.Errorf("Expected 3 scopes, got %d", len(receivedConsentVersion.Scopes))
		}

		if receivedConsentVersion.ConsentLocale.Locale != expectedConsentVersion.ConsentLocale.Locale {
			t.Errorf("Expected ConsentLocale.Locale %s, got %s", expectedConsentVersion.ConsentLocale.Locale, receivedConsentVersion.ConsentLocale.Locale)
		}

		response := ConsentVersionResponse{
			Success: true,
			Status:  201,
			Data:    expectedConsentVersion,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentVersion := NewConsentVersion(config)

	result, err := consentVersion.Upsert(context.Background(), expectedConsentVersion)

	if err != nil {
		t.Fatalf("Upsert failed: %v", err)
	}

	if !result.Success {
		t.Error("Expected success to be true")
	}

	if result.Data.ConsentID != expectedConsentVersion.ConsentID {
		t.Errorf("Expected ConsentID %s, got %s", expectedConsentVersion.ConsentID, result.Data.ConsentID)
	}

	if result.Data.Version != expectedConsentVersion.Version {
		t.Errorf("Expected Version %f, got %f", expectedConsentVersion.Version, result.Data.Version)
	}

	if len(result.Data.Scopes) != 3 {
		t.Errorf("Expected 3 scopes in response, got %d", len(result.Data.Scopes))
	}
}

func TestConsentVersion_Upsert_Error(t *testing.T) {
	server := NewMockServer(http.StatusBadRequest, `{"error": "invalid consent version"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentVersion := NewConsentVersion(config)

	consentVersionModel := ConsentVersionModel{ConsentID: "invalid"}
	_, err := consentVersion.Upsert(context.Background(), consentVersionModel)

	if err == nil {
		t.Error("Expected error for bad request, got nil")
	}
}

func TestConsentVersion_Upsert_MinimalData(t *testing.T) {
	minimalConsentVersion := ConsentVersionModel{
		ConsentID:   "consent-minimal",
		Version:     1.0,
		ConsentType: "BASIC",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var receivedConsentVersion ConsentVersionModel
		json.Unmarshal(body, &receivedConsentVersion)

		if receivedConsentVersion.ConsentID != minimalConsentVersion.ConsentID {
			t.Errorf("Expected ConsentID %s, got %s", minimalConsentVersion.ConsentID, receivedConsentVersion.ConsentID)
		}

		if receivedConsentVersion.Version != minimalConsentVersion.Version {
			t.Errorf("Expected Version %f, got %f", minimalConsentVersion.Version, receivedConsentVersion.Version)
		}

		response := ConsentVersionResponse{
			Success: true,
			Status:  201,
			Data:    minimalConsentVersion,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentVersion := NewConsentVersion(config)

	result, err := consentVersion.Upsert(context.Background(), minimalConsentVersion)

	if err != nil {
		t.Fatalf("Upsert with minimal data failed: %v", err)
	}

	if result.Data.ConsentID != minimalConsentVersion.ConsentID {
		t.Errorf("Expected ConsentID %s, got %s", minimalConsentVersion.ConsentID, result.Data.ConsentID)
	}
}

func TestConsentVersion_Upsert_WithComplexLocale(t *testing.T) {
	complexConsentVersion := ConsentVersionModel{
		ConsentID:      "consent-complex",
		Version:        2.1,
		ConsentType:    "MARKETING",
		Scopes:         []string{"marketing", "analytics", "personalization", "advertising"},
		RequiredFields: []string{"email", "given_name", "family_name", "phone_number", "address"},
		ConsentLocale: ConsentLocale{
			Locale:  "de-DE",
			Content: "Dies ist der Inhalt der Marketing-Einverständniserklärung auf Deutsch mit Sonderzeichen: äöüß",
			URL:     "https://example.com/marketing-de?param=value&other=test",
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var receivedConsentVersion ConsentVersionModel
		json.Unmarshal(body, &receivedConsentVersion)

		if len(receivedConsentVersion.Scopes) != 4 {
			t.Errorf("Expected 4 scopes, got %d", len(receivedConsentVersion.Scopes))
		}

		if len(receivedConsentVersion.RequiredFields) != 5 {
			t.Errorf("Expected 5 required fields, got %d", len(receivedConsentVersion.RequiredFields))
		}

		if receivedConsentVersion.ConsentLocale.Locale != "de-DE" {
			t.Errorf("Expected ConsentLocale.Locale 'de-DE', got %s", receivedConsentVersion.ConsentLocale.Locale)
		}

		response := ConsentVersionResponse{
			Success: true,
			Status:  201,
			Data:    complexConsentVersion,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentVersion := NewConsentVersion(config)

	result, err := consentVersion.Upsert(context.Background(), complexConsentVersion)

	if err != nil {
		t.Fatalf("Upsert with complex locale failed: %v", err)
	}

	if len(result.Data.Scopes) != 4 {
		t.Errorf("Expected 4 scopes in response, got %d", len(result.Data.Scopes))
	}

	if result.Data.ConsentLocale.Locale != "de-DE" {
		t.Errorf("Expected ConsentLocale.Locale 'de-DE', got %s", result.Data.ConsentLocale.Locale)
	}
}

func TestConsentVersion_Upsert_InvalidJSONResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{invalid json}`))
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentVersion := NewConsentVersion(config)

	consentVersionModel := ConsentVersionModel{ConsentID: "test", Version: 1.0}
	_, err := consentVersion.Upsert(context.Background(), consentVersionModel)

	if err == nil {
		t.Error("Expected error for invalid JSON response, got nil")
	}
}

func TestConsentVersion_Upsert_ContextCancellation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate slow response
		select {
		case <-r.Context().Done():
			return
		}
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentVersion := NewConsentVersion(config)

	// Create cancelled context
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	consentVersionModel := ConsentVersionModel{ConsentID: "test", Version: 1.0}
	_, err := consentVersion.Upsert(ctx, consentVersionModel)

	if err == nil {
		t.Error("Expected context cancellation error, got nil")
	}
}

func TestConsentVersion_Upsert_ServerError(t *testing.T) {
	server := NewMockServer(http.StatusInternalServerError, `{"error": "internal server error"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentVersion := NewConsentVersion(config)

	consentVersionModel := ConsentVersionModel{ConsentID: "test", Version: 1.0}
	_, err := consentVersion.Upsert(context.Background(), consentVersionModel)

	if err == nil {
		t.Error("Expected error for server error, got nil")
	}
}

func TestConsentVersion_Get_Success(t *testing.T) {
	expectedConsentVersions := []ConsentVersionModel{
		{
			ID:             "cv-1",
			Version:        1.0,
			ConsentID:      "consent-123",
			ConsentType:    "PRIVACY",
			Scopes:         []string{"profile", "email"},
			RequiredFields: []string{"email", "given_name"},
			ConsentLocale: ConsentLocale{
				Locale:  "en-US",
				Content: "Privacy consent content",
				URL:     "https://example.com/privacy-en",
			},
			CreatedAt: "2024-01-10T09:00:00Z",
			UpdatedAt: "2024-01-10T09:30:00Z",
		},
		{
			ID:             "cv-2",
			Version:        2.0,
			ConsentID:      "consent-123",
			ConsentType:    "MARKETING",
			Scopes:         []string{"marketing", "analytics"},
			RequiredFields: []string{"email", "phone_number"},
			ConsentLocale: ConsentLocale{
				Locale:  "en-US",
				Content: "Marketing consent content",
				URL:     "https://example.com/marketing-en",
			},
			CreatedAt: "2024-01-11T10:00:00Z",
			UpdatedAt: "2024-01-11T10:30:00Z",
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, "consent-management-srv/v2/consent/versions/list/consent-123") {
			t.Errorf("Expected consent-management-srv/v2/consent/versions/list/consent-123 endpoint, got %s", r.URL.Path)
		}

		// Verify Authorization header
		authHeader := r.Header.Get("Authorization")
		if !strings.Contains(authHeader, "test-token") {
			t.Errorf("Expected Authorization header with token, got %s", authHeader)
		}

		response := ConsentVersionReadResponse{
			Success: true,
			Status:  200,
			Data:    expectedConsentVersions,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentVersion := NewConsentVersion(config)

	result, err := consentVersion.Get(context.Background(), "consent-123")

	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}

	if !result.Success {
		t.Error("Expected success to be true")
	}

	if len(result.Data) != 2 {
		t.Errorf("Expected 2 consent versions, got %d", len(result.Data))
	}

	if result.Data[0].Version != 1.0 {
		t.Errorf("Expected first version 1.0, got %f", result.Data[0].Version)
	}

	if result.Data[1].Version != 2.0 {
		t.Errorf("Expected second version 2.0, got %f", result.Data[1].Version)
	}
}

func TestConsentVersion_Get_EmptyResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := ConsentVersionReadResponse{
			Success: true,
			Status:  200,
			Data:    []ConsentVersionModel{},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentVersion := NewConsentVersion(config)

	result, err := consentVersion.Get(context.Background(), "empty-consent")

	if err != nil {
		t.Fatalf("Get with empty response failed: %v", err)
	}

	if len(result.Data) != 0 {
		t.Errorf("Expected 0 consent versions, got %d", len(result.Data))
	}
}

func TestConsentVersion_Get_Error(t *testing.T) {
	server := NewMockServer(http.StatusNotFound, `{"error": "consent not found"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentVersion := NewConsentVersion(config)

	_, err := consentVersion.Get(context.Background(), "nonexistent")

	if err == nil {
		t.Error("Expected error for not found consent, got nil")
	}
}

func TestConsentVersion_Get_WithSpecialCharacters(t *testing.T) {
	consentIDWithSpecialChars := "consent_123-test.consent@domain"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify URL path contains the special character consent ID
		if !strings.Contains(r.URL.Path, consentIDWithSpecialChars) {
			t.Errorf("Expected '%s' in URL path, got %s", consentIDWithSpecialChars, r.URL.Path)
		}

		response := ConsentVersionReadResponse{
			Success: true,
			Status:  200,
			Data: []ConsentVersionModel{
				{
					ID:        "cv-special",
					ConsentID: consentIDWithSpecialChars,
					Version:   1.0,
				},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentVersion := NewConsentVersion(config)

	result, err := consentVersion.Get(context.Background(), consentIDWithSpecialChars)

	if err != nil {
		t.Fatalf("Get with special characters failed: %v", err)
	}

	if result.Data[0].ConsentID != consentIDWithSpecialChars {
		t.Errorf("Expected ConsentID %s, got %s", consentIDWithSpecialChars, result.Data[0].ConsentID)
	}
}

func TestConsentVersion_UpsertLocal_Success(t *testing.T) {
	expectedConsentLocal := ConsentLocalModel{
		ConsentVersionID: "cv-123",
		ConsentID:        "consent-456",
		Content:          "This is the localized consent content with special characters: äöüß",
		Locale:           "de-DE",
		URL:              "https://example.com/consent-de?param=value",
		Scopes:           []string{"profile", "email", "marketing"},
		RequiredFields:   []string{"email", "given_name", "family_name"},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, "consent-management-srv/v2/consent/locale") {
			t.Errorf("Expected consent-management-srv/v2/consent/locale endpoint, got %s", r.URL.Path)
		}

		// Verify Authorization header
		authHeader := r.Header.Get("Authorization")
		if !strings.Contains(authHeader, "test-token") {
			t.Errorf("Expected Authorization header with token, got %s", authHeader)
		}

		body, _ := io.ReadAll(r.Body)
		var receivedConsentLocal ConsentLocalModel
		json.Unmarshal(body, &receivedConsentLocal)

		if receivedConsentLocal.ConsentVersionID != expectedConsentLocal.ConsentVersionID {
			t.Errorf("Expected ConsentVersionID %s, got %s", expectedConsentLocal.ConsentVersionID, receivedConsentLocal.ConsentVersionID)
		}

		if receivedConsentLocal.Locale != expectedConsentLocal.Locale {
			t.Errorf("Expected Locale %s, got %s", expectedConsentLocal.Locale, receivedConsentLocal.Locale)
		}

		if len(receivedConsentLocal.Scopes) != 3 {
			t.Errorf("Expected 3 scopes, got %d", len(receivedConsentLocal.Scopes))
		}

		response := ConsentLocalResponse{
			Success: true,
			Status:  201,
			Data:    expectedConsentLocal,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentVersion := NewConsentVersion(config)

	result, err := consentVersion.UpsertLocal(context.Background(), expectedConsentLocal)

	if err != nil {
		t.Fatalf("UpsertLocal failed: %v", err)
	}

	if !result.Success {
		t.Error("Expected success to be true")
	}

	if result.Data.ConsentVersionID != expectedConsentLocal.ConsentVersionID {
		t.Errorf("Expected ConsentVersionID %s, got %s", expectedConsentLocal.ConsentVersionID, result.Data.ConsentVersionID)
	}

	if result.Data.Locale != expectedConsentLocal.Locale {
		t.Errorf("Expected Locale %s, got %s", expectedConsentLocal.Locale, result.Data.Locale)
	}
}

func TestConsentVersion_UpsertLocal_Error(t *testing.T) {
	server := NewMockServer(http.StatusBadRequest, `{"error": "invalid consent local"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentVersion := NewConsentVersion(config)

	consentLocalModel := ConsentLocalModel{ConsentVersionID: "invalid"}
	_, err := consentVersion.UpsertLocal(context.Background(), consentLocalModel)

	if err == nil {
		t.Error("Expected error for bad request, got nil")
	}
}

func TestConsentVersion_UpsertLocal_MinimalData(t *testing.T) {
	minimalConsentLocal := ConsentLocalModel{
		ConsentVersionID: "cv-minimal",
		Locale:           "en",
		Content:          "Minimal consent content",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var receivedConsentLocal ConsentLocalModel
		json.Unmarshal(body, &receivedConsentLocal)

		if receivedConsentLocal.ConsentVersionID != minimalConsentLocal.ConsentVersionID {
			t.Errorf("Expected ConsentVersionID %s, got %s", minimalConsentLocal.ConsentVersionID, receivedConsentLocal.ConsentVersionID)
		}

		response := ConsentLocalResponse{
			Success: true,
			Status:  201,
			Data:    minimalConsentLocal,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentVersion := NewConsentVersion(config)

	result, err := consentVersion.UpsertLocal(context.Background(), minimalConsentLocal)

	if err != nil {
		t.Fatalf("UpsertLocal with minimal data failed: %v", err)
	}

	if result.Data.ConsentVersionID != minimalConsentLocal.ConsentVersionID {
		t.Errorf("Expected ConsentVersionID %s, got %s", minimalConsentLocal.ConsentVersionID, result.Data.ConsentVersionID)
	}
}

func TestConsentVersion_Get_InvalidJSONResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{invalid json}`))
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentVersion := NewConsentVersion(config)

	_, err := consentVersion.Get(context.Background(), "test-consent")

	if err == nil {
		t.Error("Expected error for invalid JSON response, got nil")
	}
}

func TestConsentVersion_UpsertLocal_InvalidJSONResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{invalid json}`))
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentVersion := NewConsentVersion(config)

	consentLocalModel := ConsentLocalModel{ConsentVersionID: "test", Locale: "en"}
	_, err := consentVersion.UpsertLocal(context.Background(), consentLocalModel)

	if err == nil {
		t.Error("Expected error for invalid JSON response, got nil")
	}
}

func TestConsentVersion_GetLocal_Success(t *testing.T) {
	expectedConsentLocal := ConsentLocalModel{
		ConsentVersionID: "cv-123",
		ConsentID:        "consent-456",
		Content:          "This is the German consent content with special characters: äöüß",
		Locale:           "de-DE",
		URL:              "https://example.com/consent-de?param=value&other=test",
		Scopes:           []string{"profile", "email", "marketing"},
		RequiredFields:   []string{"email", "given_name", "family_name"},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, "consent-management-srv/v2/consent/locale/cv-123") {
			t.Errorf("Expected consent-management-srv/v2/consent/locale/cv-123 in path, got %s", r.URL.Path)
		}

		// Verify locale query parameter
		localeParam := r.URL.Query().Get("locale")
		if localeParam != "de-DE" {
			t.Errorf("Expected locale parameter 'de-DE', got %s", localeParam)
		}

		// Verify Authorization header
		authHeader := r.Header.Get("Authorization")
		if !strings.Contains(authHeader, "test-token") {
			t.Errorf("Expected Authorization header with token, got %s", authHeader)
		}

		response := ConsentLocalResponse{
			Success: true,
			Status:  200,
			Data:    expectedConsentLocal,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentVersion := NewConsentVersion(config)

	result, err := consentVersion.GetLocal(context.Background(), "cv-123", "de-DE")

	if err != nil {
		t.Fatalf("GetLocal failed: %v", err)
	}

	if !result.Success {
		t.Error("Expected success to be true")
	}

	if result.Data.ConsentVersionID != expectedConsentLocal.ConsentVersionID {
		t.Errorf("Expected ConsentVersionID %s, got %s", expectedConsentLocal.ConsentVersionID, result.Data.ConsentVersionID)
	}

	if result.Data.Locale != expectedConsentLocal.Locale {
		t.Errorf("Expected Locale %s, got %s", expectedConsentLocal.Locale, result.Data.Locale)
	}

	if len(result.Data.Scopes) != 3 {
		t.Errorf("Expected 3 scopes, got %d", len(result.Data.Scopes))
	}
}

func TestConsentVersion_GetLocal_NoContent(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentVersion := NewConsentVersion(config)

	result, err := consentVersion.GetLocal(context.Background(), "cv-empty", "en-US")

	if err != nil {
		t.Fatalf("GetLocal with no content failed: %v", err)
	}

	if result.Success {
		t.Error("Expected success to be false for no content")
	}

	if result.Status != http.StatusNoContent {
		t.Errorf("Expected status %d, got %d", http.StatusNoContent, result.Status)
	}
}

func TestConsentVersion_GetLocal_Error(t *testing.T) {
	server := NewMockServer(http.StatusNotFound, `{"error": "consent local not found"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentVersion := NewConsentVersion(config)

	_, err := consentVersion.GetLocal(context.Background(), "nonexistent", "en-US")

	if err == nil {
		t.Error("Expected error for not found consent local, got nil")
	}
}

func TestConsentVersion_GetLocal_WithSpecialCharacters(t *testing.T) {
	consentVersionIDWithSpecialChars := "cv_123-test.version@domain"
	localeWithSpecialChars := "zh-CN"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify URL path contains the special character consent version ID
		if !strings.Contains(r.URL.Path, consentVersionIDWithSpecialChars) {
			t.Errorf("Expected '%s' in URL path, got %s", consentVersionIDWithSpecialChars, r.URL.Path)
		}

		// Verify locale query parameter
		localeParam := r.URL.Query().Get("locale")
		if localeParam != localeWithSpecialChars {
			t.Errorf("Expected locale parameter '%s', got %s", localeWithSpecialChars, localeParam)
		}

		response := ConsentLocalResponse{
			Success: true,
			Status:  200,
			Data: ConsentLocalModel{
				ConsentVersionID: consentVersionIDWithSpecialChars,
				Locale:           localeWithSpecialChars,
				Content:          "Chinese consent content: 这是中文同意书内容",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentVersion := NewConsentVersion(config)

	result, err := consentVersion.GetLocal(context.Background(), consentVersionIDWithSpecialChars, localeWithSpecialChars)

	if err != nil {
		t.Fatalf("GetLocal with special characters failed: %v", err)
	}

	if result.Data.ConsentVersionID != consentVersionIDWithSpecialChars {
		t.Errorf("Expected ConsentVersionID %s, got %s", consentVersionIDWithSpecialChars, result.Data.ConsentVersionID)
	}

	if result.Data.Locale != localeWithSpecialChars {
		t.Errorf("Expected Locale %s, got %s", localeWithSpecialChars, result.Data.Locale)
	}
}

func TestConsentVersion_GetLocal_EmptyParameters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify URL path ends with empty parameters
		if !strings.HasSuffix(r.URL.Path, "consent-management-srv/v2/consent/locale/") {
			t.Errorf("Expected URL to end with 'consent-management-srv/v2/consent/locale/', got %s", r.URL.Path)
		}

		// Verify empty locale query parameter
		localeParam := r.URL.Query().Get("locale")
		if localeParam != "" {
			t.Errorf("Expected empty locale parameter, got %s", localeParam)
		}

		response := ConsentLocalResponse{
			Success: true,
			Status:  200,
			Data: ConsentLocalModel{
				ConsentVersionID: "",
				Locale:           "",
				Content:          "Empty parameters consent",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentVersion := NewConsentVersion(config)

	result, err := consentVersion.GetLocal(context.Background(), "", "")

	if err != nil {
		t.Fatalf("GetLocal with empty parameters failed: %v", err)
	}

	if result.Data.Content != "Empty parameters consent" {
		t.Errorf("Expected Content 'Empty parameters consent', got %s", result.Data.Content)
	}
}

func TestConsentVersion_GetLocal_InvalidJSONResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{invalid json}`))
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentVersion := NewConsentVersion(config)

	_, err := consentVersion.GetLocal(context.Background(), "test-cv", "en-US")

	if err == nil {
		t.Error("Expected error for invalid JSON response, got nil")
	}
}

// func TestConsentVersion_GetLocal_ContextCancellation(t *testing.T) {
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Simulate slow response
// 		select {
// 		case <-r.Context().Done():
// 			return
// 		}
// 	}))
// 	defer server.Close()

// 	config := NewTestClientConfig(server.URL)
// 	consentVersion := NewConsentVersion(config)

// 	// Create cancelled context
// 	ctx, cancel := context.WithCancel(context.Background())
// 	cancel()

// 	_, err := consentVersion.GetLocal(ctx, "test-cv", "en-US")

// 	if err == nil {
// 		t.Error("Expected context cancellation error, got nil")
// 	}
// }

func TestConsentVersion_GetLocal_ServerError(t *testing.T) {
	server := NewMockServer(http.StatusInternalServerError, `{"error": "internal server error"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentVersion := NewConsentVersion(config)

	_, err := consentVersion.GetLocal(context.Background(), "test-cv", "en-US")

	if err == nil {
		t.Error("Expected error for server error, got nil")
	}
}

func TestConsentVersion_GetLocal_MultipleLocales(t *testing.T) {
	locales := []string{"en-US", "de-DE", "fr-FR", "es-ES"}
	requestedLocales := make(map[string]bool)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		localeParam := r.URL.Query().Get("locale")
		requestedLocales[localeParam] = true

		response := ConsentLocalResponse{
			Success: true,
			Status:  200,
			Data: ConsentLocalModel{
				ConsentVersionID: "cv-multi",
				Locale:           localeParam,
				Content:          fmt.Sprintf("Content for locale %s", localeParam),
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	consentVersion := NewConsentVersion(config)

	// Request multiple locales
	for _, locale := range locales {
		result, err := consentVersion.GetLocal(context.Background(), "cv-multi", locale)
		if err != nil {
			t.Fatalf("GetLocal failed for locale %s: %v", locale, err)
		}

		if result.Data.Locale != locale {
			t.Errorf("Expected locale %s, got %s", locale, result.Data.Locale)
		}
	}

	// Verify all locales were processed
	for _, locale := range locales {
		if !requestedLocales[locale] {
			t.Errorf("Locale %s was not requested", locale)
		}
	}
}
