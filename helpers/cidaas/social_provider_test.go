// helpers/cidaas/social_provider_test.go
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

func TestNewSocialProvider(t *testing.T) {
	config := NewTestClientConfig("http://test.com")

	socialProvider := NewSocialProvider(config)

	if socialProvider == nil {
		t.Fatal("Expected social provider instance, got nil")
	}

	if socialProvider.BaseURL != config.BaseURL {
		t.Errorf("Expected BaseURL %s, got %s", config.BaseURL, socialProvider.BaseURL)
	}

	if socialProvider.AccessToken != config.AccessToken {
		t.Errorf("Expected AccessToken %s, got %s", config.AccessToken, socialProvider.AccessToken)
	}
}

func TestSocialProvider_Upsert_Success(t *testing.T) {
	expectedProvider := &SocialProviderModel{
		ID:                    "provider-123",
		ClientID:              "google-client-123",
		ClientSecret:          "google-secret-456",
		Name:                  "Google OAuth",
		ProviderName:          "google",
		EnabledForAdminPortal: true,
		Enabled:               true,
		Scopes:                []string{"openid", "profile", "email"},
		Claims: &ClaimsModel{
			RequiredClaims: RequiredClaimsModel{
				UserInfo: []string{"email", "name"},
				IDToken:  []string{"sub", "email"},
			},
			OptionalClaims: OptionalClaimsModel{
				UserInfo: []string{"picture", "locale"},
				IDToken:  []string{"given_name", "family_name"},
			},
		},
		UserInfoFields: []UserInfoFieldsModel{
			{
				InnerKey:      "email",
				ExternalKey:   "email",
				IsCustomField: false,
				IsSystemField: true,
			},
			{
				InnerKey:      "given_name",
				ExternalKey:   "first_name",
				IsCustomField: false,
				IsSystemField: true,
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method and endpoint
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, "providers-srv/multi/providers") {
			t.Errorf("Expected providers-srv/multi/providers endpoint, got %s", r.URL.Path)
		}

		// Verify request body
		body, _ := io.ReadAll(r.Body)
		var receivedProvider SocialProviderModel
		json.Unmarshal(body, &receivedProvider)

		if receivedProvider.ProviderName != expectedProvider.ProviderName {
			t.Errorf("Expected ProviderName %s, got %s", expectedProvider.ProviderName, receivedProvider.ProviderName)
		}

		if receivedProvider.ClientID != expectedProvider.ClientID {
			t.Errorf("Expected ClientID %s, got %s", expectedProvider.ClientID, receivedProvider.ClientID)
		}

		if receivedProvider.Enabled != expectedProvider.Enabled {
			t.Errorf("Expected Enabled %t, got %t", expectedProvider.Enabled, receivedProvider.Enabled)
		}

		if len(receivedProvider.Scopes) != len(expectedProvider.Scopes) {
			t.Errorf("Expected %d scopes, got %d", len(expectedProvider.Scopes), len(receivedProvider.Scopes))
		}

		// Verify Claims structure
		if receivedProvider.Claims == nil {
			t.Error("Expected Claims to be present")
		} else {
			if len(receivedProvider.Claims.RequiredClaims.UserInfo) != 2 {
				t.Errorf("Expected 2 required user info claims, got %d", len(receivedProvider.Claims.RequiredClaims.UserInfo))
			}
		}

		// Verify UserInfoFields
		if len(receivedProvider.UserInfoFields) != 2 {
			t.Errorf("Expected 2 user info fields, got %d", len(receivedProvider.UserInfoFields))
		}

		// Send success response
		response := SocialProviderResponse{
			Success: true,
			Status:  200,
			Data:    *expectedProvider,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	result, err := socialProvider.Upsert(context.Background(), expectedProvider)

	if err != nil {
		t.Fatalf("Upsert failed: %v", err)
	}

	if !result.Success {
		t.Error("Expected success to be true")
	}

	if result.Data.ProviderName != expectedProvider.ProviderName {
		t.Errorf("Expected ProviderName %s, got %s", expectedProvider.ProviderName, result.Data.ProviderName)
	}

	if result.Data.ClientID != expectedProvider.ClientID {
		t.Errorf("Expected ClientID %s, got %s", expectedProvider.ClientID, result.Data.ClientID)
	}

	if len(result.Data.Scopes) != len(expectedProvider.Scopes) {
		t.Errorf("Expected %d scopes, got %d", len(expectedProvider.Scopes), len(result.Data.Scopes))
	}
}

func TestSocialProvider_Upsert_MinimalData(t *testing.T) {
	minimalProvider := &SocialProviderModel{
		ProviderName: "facebook",
		ClientID:     "fb-client-123",
		Enabled:      true,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var receivedProvider SocialProviderModel
		json.Unmarshal(body, &receivedProvider)

		if receivedProvider.ProviderName != minimalProvider.ProviderName {
			t.Errorf("Expected ProviderName %s, got %s", minimalProvider.ProviderName, receivedProvider.ProviderName)
		}

		if receivedProvider.ClientID != minimalProvider.ClientID {
			t.Errorf("Expected ClientID %s, got %s", minimalProvider.ClientID, receivedProvider.ClientID)
		}

		response := SocialProviderResponse{
			Success: true,
			Status:  200,
			Data:    *minimalProvider,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	result, err := socialProvider.Upsert(context.Background(), minimalProvider)

	if err != nil {
		t.Fatalf("Upsert with minimal data failed: %v", err)
	}

	if !result.Success {
		t.Error("Expected success to be true")
	}

	if result.Data.ProviderName != minimalProvider.ProviderName {
		t.Errorf("Expected ProviderName %s, got %s", minimalProvider.ProviderName, result.Data.ProviderName)
	}
}

func TestSocialProvider_Upsert_WithComplexClaims(t *testing.T) {
	providerWithClaims := &SocialProviderModel{
		ProviderName: "microsoft",
		ClientID:     "ms-client-123",
		Enabled:      true,
		Claims: &ClaimsModel{
			RequiredClaims: RequiredClaimsModel{
				UserInfo: []string{"email", "name", "preferred_username"},
				IDToken:  []string{"sub", "email", "name"},
			},
			OptionalClaims: OptionalClaimsModel{
				UserInfo: []string{"picture", "locale", "zoneinfo"},
				IDToken:  []string{"given_name", "family_name", "middle_name"},
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var receivedProvider SocialProviderModel
		json.Unmarshal(body, &receivedProvider)

		// Verify complex claims structure
		if receivedProvider.Claims == nil {
			t.Fatal("Expected Claims to be present")
		}

		if len(receivedProvider.Claims.RequiredClaims.UserInfo) != 3 {
			t.Errorf("Expected 3 required user info claims, got %d", len(receivedProvider.Claims.RequiredClaims.UserInfo))
		}

		if len(receivedProvider.Claims.OptionalClaims.IDToken) != 3 {
			t.Errorf("Expected 3 optional ID token claims, got %d", len(receivedProvider.Claims.OptionalClaims.IDToken))
		}

		response := SocialProviderResponse{
			Success: true,
			Status:  200,
			Data:    *providerWithClaims,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	result, err := socialProvider.Upsert(context.Background(), providerWithClaims)

	if err != nil {
		t.Fatalf("Upsert with complex claims failed: %v", err)
	}

	if result.Data.Claims == nil {
		t.Error("Expected Claims in response")
	}

	if len(result.Data.Claims.RequiredClaims.UserInfo) != 3 {
		t.Errorf("Expected 3 required user info claims in response, got %d", len(result.Data.Claims.RequiredClaims.UserInfo))
	}
}

func TestSocialProvider_Upsert_WithUserInfoFields(t *testing.T) {
	providerWithFields := &SocialProviderModel{
		ProviderName: "github",
		ClientID:     "gh-client-123",
		Enabled:      true,
		UserInfoFields: []UserInfoFieldsModel{
			{
				InnerKey:      "email",
				ExternalKey:   "email",
				IsCustomField: false,
				IsSystemField: true,
			},
			{
				InnerKey:      "username",
				ExternalKey:   "login",
				IsCustomField: true,
				IsSystemField: false,
			},
			{
				InnerKey:      "avatar",
				ExternalKey:   "avatar_url",
				IsCustomField: true,
				IsSystemField: false,
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var receivedProvider SocialProviderModel
		json.Unmarshal(body, &receivedProvider)

		if len(receivedProvider.UserInfoFields) != 3 {
			t.Errorf("Expected 3 user info fields, got %d", len(receivedProvider.UserInfoFields))
		}

		// Verify specific field mappings
		emailField := receivedProvider.UserInfoFields[0]
		if emailField.InnerKey != "email" || emailField.ExternalKey != "email" {
			t.Errorf("Expected email field mapping, got inner: %s, external: %s", emailField.InnerKey, emailField.ExternalKey)
		}

		if !emailField.IsSystemField || emailField.IsCustomField {
			t.Error("Expected email field to be system field, not custom field")
		}

		response := SocialProviderResponse{
			Success: true,
			Status:  200,
			Data:    *providerWithFields,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	result, err := socialProvider.Upsert(context.Background(), providerWithFields)

	if err != nil {
		t.Fatalf("Upsert with user info fields failed: %v", err)
	}

	if len(result.Data.UserInfoFields) != 3 {
		t.Errorf("Expected 3 user info fields in response, got %d", len(result.Data.UserInfoFields))
	}
}

func TestSocialProvider_Upsert_ServerError(t *testing.T) {
	server := NewMockServer(http.StatusInternalServerError, `{"error": "internal server error"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	providerModel := &SocialProviderModel{ProviderName: "test"}
	_, err := socialProvider.Upsert(context.Background(), providerModel)

	if err == nil {
		t.Error("Expected error for server error, got nil")
	}
}

func TestSocialProvider_Upsert_InvalidJSON(t *testing.T) {
	server := NewMockServer(http.StatusOK, `{invalid json}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	providerModel := &SocialProviderModel{ProviderName: "test"}
	_, err := socialProvider.Upsert(context.Background(), providerModel)

	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}

func TestSocialProvider_Upsert_ContextCancellation(t *testing.T) {
	server := NewMockServer(http.StatusOK, `{"success": true}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	// Create cancelled context
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	providerModel := &SocialProviderModel{ProviderName: "test"}
	_, err := socialProvider.Upsert(ctx, providerModel)

	if err == nil {
		t.Error("Expected context cancellation error, got nil")
	}
}

func TestSocialProvider_Upsert_DisabledProvider(t *testing.T) {
	disabledProvider := &SocialProviderModel{
		ProviderName:          "twitter",
		ClientID:              "twitter-client-123",
		Enabled:               false, // Disabled provider
		EnabledForAdminPortal: false,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var receivedProvider SocialProviderModel
		json.Unmarshal(body, &receivedProvider)

		if receivedProvider.Enabled {
			t.Error("Expected provider to be disabled")
		}

		if receivedProvider.EnabledForAdminPortal {
			t.Error("Expected provider to be disabled for admin portal")
		}

		response := SocialProviderResponse{
			Success: true,
			Status:  200,
			Data:    *disabledProvider,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	result, err := socialProvider.Upsert(context.Background(), disabledProvider)

	if err != nil {
		t.Fatalf("Upsert disabled provider failed: %v", err)
	}

	if result.Data.Enabled {
		t.Error("Expected provider to be disabled in response")
	}

	if result.Data.EnabledForAdminPortal {
		t.Error("Expected provider to be disabled for admin portal in response")
	}
}

func TestSocialProvider_Upsert_WithMultipleScopes(t *testing.T) {
	providerWithScopes := &SocialProviderModel{
		ProviderName: "linkedin",
		ClientID:     "linkedin-client-123",
		Enabled:      true,
		Scopes:       []string{"r_liteprofile", "r_emailaddress", "w_member_social", "r_basicprofile"},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var receivedProvider SocialProviderModel
		json.Unmarshal(body, &receivedProvider)

		if len(receivedProvider.Scopes) != 4 {
			t.Errorf("Expected 4 scopes, got %d", len(receivedProvider.Scopes))
		}

		expectedScopes := []string{"r_liteprofile", "r_emailaddress", "w_member_social", "r_basicprofile"}
		for i, scope := range receivedProvider.Scopes {
			if scope != expectedScopes[i] {
				t.Errorf("Expected scope %s at index %d, got %s", expectedScopes[i], i, scope)
			}
		}

		response := SocialProviderResponse{
			Success: true,
			Status:  200,
			Data:    *providerWithScopes,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	result, err := socialProvider.Upsert(context.Background(), providerWithScopes)

	if err != nil {
		t.Fatalf("Upsert with multiple scopes failed: %v", err)
	}

	if len(result.Data.Scopes) != 4 {
		t.Errorf("Expected 4 scopes in response, got %d", len(result.Data.Scopes))
	}
}

// Add these tests to helpers/cidaas/social_provider_test.go

func TestSocialProvider_Get_Success(t *testing.T) {
	expectedProvider := SocialProviderModel{
		ID:                    "provider-123",
		ClientID:              "google-client-123",
		ClientSecret:          "google-secret-456",
		Name:                  "Google OAuth",
		ProviderName:          "google",
		EnabledForAdminPortal: true,
		Enabled:               true,
		Scopes:                []string{"openid", "profile", "email"},
		Claims: &ClaimsModel{
			RequiredClaims: RequiredClaimsModel{
				UserInfo: []string{"email", "name"},
				IDToken:  []string{"sub", "email"},
			},
		},
		UserInfoFields: []UserInfoFieldsModel{
			{
				InnerKey:      "email",
				ExternalKey:   "email",
				IsSystemField: true,
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method and endpoint
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, "providers-srv/multi/providers") {
			t.Errorf("Expected providers-srv/multi/providers endpoint, got %s", r.URL.Path)
		}

		// Verify query parameters
		providerNameParam := r.URL.Query().Get("provider_name")
		if providerNameParam != "google" {
			t.Errorf("Expected provider_name parameter 'google', got %s", providerNameParam)
		}

		providerIDParam := r.URL.Query().Get("provider_id")
		if providerIDParam != "provider-123" {
			t.Errorf("Expected provider_id parameter 'provider-123', got %s", providerIDParam)
		}

		// Send success response
		response := SocialProviderResponse{
			Success: true,
			Status:  200,
			Data:    expectedProvider,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	result, err := socialProvider.Get(context.Background(), "google", "provider-123")

	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}

	if !result.Success {
		t.Error("Expected success to be true")
	}

	if result.Data.ProviderName != expectedProvider.ProviderName {
		t.Errorf("Expected ProviderName %s, got %s", expectedProvider.ProviderName, result.Data.ProviderName)
	}

	if result.Data.ClientID != expectedProvider.ClientID {
		t.Errorf("Expected ClientID %s, got %s", expectedProvider.ClientID, result.Data.ClientID)
	}

	if len(result.Data.Scopes) != len(expectedProvider.Scopes) {
		t.Errorf("Expected %d scopes, got %d", len(expectedProvider.Scopes), len(result.Data.Scopes))
	}
}

func TestSocialProvider_Get_WithSpecialCharacters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify query parameters with special characters
		providerNameParam := r.URL.Query().Get("provider_name")
		if providerNameParam != "custom-provider" {
			t.Errorf("Expected provider_name parameter 'custom-provider', got %s", providerNameParam)
		}

		providerIDParam := r.URL.Query().Get("provider_id")
		if providerIDParam != "provider_123" {
			t.Errorf("Expected provider_id parameter 'provider_123', got %s", providerIDParam)
		}

		response := SocialProviderResponse{
			Success: true,
			Status:  200,
			Data: SocialProviderModel{
				ProviderName: "custom-provider",
				ID:           "provider_123",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	_, err := socialProvider.Get(context.Background(), "custom-provider", "provider_123")

	if err != nil {
		t.Fatalf("Get with special characters failed: %v", err)
	}
}

func TestSocialProvider_Get_NotFound(t *testing.T) {
	server := NewMockServer(http.StatusNotFound, `{"error": "provider not found"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	_, err := socialProvider.Get(context.Background(), "nonexistent", "provider-999")

	if err == nil {
		t.Error("Expected error for not found provider, got nil")
	}
}

func TestSocialProvider_Get_ServerError(t *testing.T) {
	server := NewMockServer(http.StatusInternalServerError, `{"error": "internal server error"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	_, err := socialProvider.Get(context.Background(), "google", "provider-123")

	if err == nil {
		t.Error("Expected error for server error, got nil")
	}
}

func TestSocialProvider_Get_InvalidJSON(t *testing.T) {
	server := NewMockServer(http.StatusOK, `{invalid json}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	_, err := socialProvider.Get(context.Background(), "google", "provider-123")

	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}

func TestSocialProvider_Delete_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method and endpoint
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE method, got %s", r.Method)
		}

		// Verify URL path contains provider name and ID
		if !strings.Contains(r.URL.Path, "google") {
			t.Errorf("Expected 'google' in URL path, got %s", r.URL.Path)
		}

		if !strings.Contains(r.URL.Path, "provider-123") {
			t.Errorf("Expected 'provider-123' in URL path, got %s", r.URL.Path)
		}

		if !strings.Contains(r.URL.Path, "providers-srv/multi/providers") {
			t.Errorf("Expected providers-srv/multi/providers endpoint, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	err := socialProvider.Delete(context.Background(), "google", "provider-123")

	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
}

func TestSocialProvider_Delete_WithSpecialCharacters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify URL path contains provider name and ID with special characters
		if !strings.Contains(r.URL.Path, "custom-provider") {
			t.Errorf("Expected 'custom-provider' in URL path, got %s", r.URL.Path)
		}

		if !strings.Contains(r.URL.Path, "provider_123") {
			t.Errorf("Expected 'provider_123' in URL path, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	err := socialProvider.Delete(context.Background(), "custom-provider", "provider_123")

	if err != nil {
		t.Fatalf("Delete with special characters failed: %v", err)
	}
}

func TestSocialProvider_Delete_ServerError(t *testing.T) {
	server := NewMockServer(http.StatusInternalServerError, `{"error": "internal server error"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	err := socialProvider.Delete(context.Background(), "google", "provider-123")

	if err == nil {
		t.Error("Expected error for server error, got nil")
	}
}

func TestSocialProvider_Delete_NotFound(t *testing.T) {
	server := NewMockServer(http.StatusNotFound, `{"error": "provider not found"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	err := socialProvider.Delete(context.Background(), "nonexistent", "provider-999")

	if err == nil {
		t.Error("Expected error for not found provider, got nil")
	}
}

func TestSocialProvider_GetAll_Success(t *testing.T) {
	expectedProviders := []SocialProviderModel{
		{
			ID:           "provider-1",
			ProviderName: "google",
			ClientID:     "google-client-123",
			Enabled:      true,
			Scopes:       []string{"openid", "profile", "email"},
			Claims: &ClaimsModel{
				RequiredClaims: RequiredClaimsModel{
					UserInfo: []string{"email"},
				},
			},
		},
		{
			ID:           "provider-2",
			ProviderName: "facebook",
			ClientID:     "fb-client-456",
			Enabled:      true,
			Scopes:       []string{"public_profile", "email"},
			UserInfoFields: []UserInfoFieldsModel{
				{
					InnerKey:    "email",
					ExternalKey: "email",
				},
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method and endpoint
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, "providers-srv/providers/enabled/list") {
			t.Errorf("Expected providers-srv/providers/enabled/list endpoint, got %s", r.URL.Path)
		}

		// Send success response
		response := AllSocialProviderResponse{
			Success: true,
			Status:  200,
			Data:    expectedProviders,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	result, err := socialProvider.GetAll(context.Background())

	if err != nil {
		t.Fatalf("GetAll failed: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("Expected 2 providers, got %d", len(result))
	}

	if result[0].ProviderName != "google" {
		t.Errorf("Expected first provider name 'google', got %s", result[0].ProviderName)
	}

	if result[1].ProviderName != "facebook" {
		t.Errorf("Expected second provider name 'facebook', got %s", result[1].ProviderName)
	}

	// Verify complex nested data
	if result[0].Claims == nil {
		t.Error("Expected Claims for first provider")
	}

	if len(result[1].UserInfoFields) != 1 {
		t.Errorf("Expected 1 user info field for second provider, got %d", len(result[1].UserInfoFields))
	}
}

func TestSocialProvider_GetAll_EmptyResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := AllSocialProviderResponse{
			Success: true,
			Status:  200,
			Data:    []SocialProviderModel{},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	result, err := socialProvider.GetAll(context.Background())

	if err != nil {
		t.Fatalf("GetAll failed: %v", err)
	}

	if len(result) != 0 {
		t.Errorf("Expected 0 providers, got %d", len(result))
	}
}

func TestSocialProvider_GetAll_ServerError(t *testing.T) {
	server := NewMockServer(http.StatusInternalServerError, `{"error": "internal server error"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	_, err := socialProvider.GetAll(context.Background())

	if err == nil {
		t.Error("Expected error for server error, got nil")
	}
}

func TestSocialProvider_GetAll_InvalidJSON(t *testing.T) {
	server := NewMockServer(http.StatusOK, `{invalid json}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	_, err := socialProvider.GetAll(context.Background())

	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}

func TestSocialProvider_ContextCancellation_Get(t *testing.T) {
	server := NewMockServer(http.StatusOK, `{"success": true}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	// Create cancelled context
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := socialProvider.Get(ctx, "google", "provider-123")

	if err == nil {
		t.Error("Expected context cancellation error, got nil")
	}
}

func TestSocialProvider_ContextCancellation_Delete(t *testing.T) {
	server := NewMockServer(http.StatusOK, `{"success": true}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	// Create cancelled context
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := socialProvider.Delete(ctx, "google", "provider-123")

	if err == nil {
		t.Error("Expected context cancellation error, got nil")
	}
}

func TestSocialProvider_ContextCancellation_GetAll(t *testing.T) {
	server := NewMockServer(http.StatusOK, `{"success": true}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	socialProvider := NewSocialProvider(config)

	// Create cancelled context
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := socialProvider.GetAll(ctx)

	if err == nil {
		t.Error("Expected context cancellation error, got nil")
	}
}
