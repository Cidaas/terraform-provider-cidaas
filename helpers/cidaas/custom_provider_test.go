// helpers/cidaas/custom_provider_test.go
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

func TestNewCustomProvider(t *testing.T) {
	config := NewTestClientConfig("http://test.com")

	customProvider := NewCustomProvider(config)

	if customProvider == nil {
		t.Fatal("Expected custom provider instance, got nil")
	}

	if customProvider.BaseURL != config.BaseURL {
		t.Errorf("Expected BaseURL %s, got %s", config.BaseURL, customProvider.BaseURL)
	}

	if customProvider.AccessToken != config.AccessToken {
		t.Errorf("Expected AccessToken %s, got %s", config.AccessToken, customProvider.AccessToken)
	}
}

func TestCustomProvider_CreateCustomProvider_Success(t *testing.T) {
	expectedProvider := &CustomProviderModel{
		ID:                    "cp-123",
		ClientID:              "oauth-client-123",
		ClientSecret:          "oauth-secret-456",
		DisplayName:           "Custom OAuth Provider",
		StandardType:          "OAUTH2",
		AuthorizationEndpoint: "https://example.com/oauth/authorize",
		TokenEndpoint:         "https://example.com/oauth/token",
		ProviderName:          "custom-oauth",
		LogoURL:               "https://example.com/logo.png",
		UserinfoEndpoint:      "https://example.com/oauth/userinfo",
		UserinfoFields: map[string]interface{}{
			"email": "email",
			"name":  "full_name",
		},
		Scopes: Scopes{
			DisplayLabel: "OAuth Scopes",
			Scopes: []ScopeChild{
				{
					ScopeName:   "openid",
					Required:    true,
					Recommended: true,
				},
				{
					ScopeName:   "profile",
					Required:    false,
					Recommended: true,
				},
			},
		},
		Domains: []string{"example.com", "test.com"},
		AmrConfig: []AmrConfig{
			{
				AmrValue:    "pwd",
				ExtAmrValue: "password",
			},
			{
				AmrValue:    "mfa",
				ExtAmrValue: "multi_factor",
			},
		},
		UserInfoSource: "userinfo_endpoint",
		Pkce:           true,
		AuthType:       "oauth2",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, "providers-srv/custom") {
			t.Errorf("Expected providers-srv/custom endpoint, got %s", r.URL.Path)
		}

		body, _ := io.ReadAll(r.Body)
		var receivedProvider CustomProviderModel
		json.Unmarshal(body, &receivedProvider)

		if receivedProvider.ProviderName != expectedProvider.ProviderName {
			t.Errorf("Expected ProviderName %s, got %s", expectedProvider.ProviderName, receivedProvider.ProviderName)
		}

		if len(receivedProvider.Scopes.Scopes) != 2 {
			t.Errorf("Expected 2 scopes, got %d", len(receivedProvider.Scopes.Scopes))
		}

		if len(receivedProvider.AmrConfig) != 2 {
			t.Errorf("Expected 2 AMR configs, got %d", len(receivedProvider.AmrConfig))
		}

		response := CustomProviderResponse{
			Success: true,
			Status:  201,
			Data:    *expectedProvider,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	customProvider := NewCustomProvider(config)

	result, err := customProvider.CreateCustomProvider(context.Background(), expectedProvider)

	if err != nil {
		t.Fatalf("CreateCustomProvider failed: %v", err)
	}

	if !result.Success {
		t.Error("Expected success to be true")
	}

	if result.Data.ProviderName != expectedProvider.ProviderName {
		t.Errorf("Expected ProviderName %s, got %s", expectedProvider.ProviderName, result.Data.ProviderName)
	}

	if result.Data.Pkce != expectedProvider.Pkce {
		t.Errorf("Expected Pkce %t, got %t", expectedProvider.Pkce, result.Data.Pkce)
	}
}

func TestCustomProvider_CreateCustomProvider_Error(t *testing.T) {
	server := NewMockServer(http.StatusBadRequest, `{"error": "invalid provider"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	customProvider := NewCustomProvider(config)

	providerModel := &CustomProviderModel{ProviderName: "invalid"}
	_, err := customProvider.CreateCustomProvider(context.Background(), providerModel)

	if err == nil {
		t.Error("Expected error for bad request, got nil")
	}
}

func TestCustomProvider_UpdateCustomProvider_Success(t *testing.T) {
	updateProvider := &CustomProviderModel{
		ID:           "cp-123",
		ProviderName: "updated-provider",
		DisplayName:  "Updated Provider",
		Pkce:         false,
		Domains:      []string{"updated.com"},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, "providers-srv/custom") {
			t.Errorf("Expected providers-srv/custom endpoint, got %s", r.URL.Path)
		}

		body, _ := io.ReadAll(r.Body)
		var receivedProvider CustomProviderModel
		json.Unmarshal(body, &receivedProvider)

		if receivedProvider.ID != updateProvider.ID {
			t.Errorf("Expected ID %s, got %s", updateProvider.ID, receivedProvider.ID)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	customProvider := NewCustomProvider(config)

	err := customProvider.UpdateCustomProvider(context.Background(), updateProvider)

	if err != nil {
		t.Fatalf("UpdateCustomProvider failed: %v", err)
	}
}

func TestCustomProvider_UpdateCustomProvider_Error(t *testing.T) {
	server := NewMockServer(http.StatusInternalServerError, `{"error": "server error"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	customProvider := NewCustomProvider(config)

	updateProvider := &CustomProviderModel{ID: "cp-123"}
	err := customProvider.UpdateCustomProvider(context.Background(), updateProvider)

	if err == nil {
		t.Error("Expected error for server error, got nil")
	}
}

func TestCustomProvider_GetCustomProvider_Success(t *testing.T) {
	expectedProvider := CustomProviderModel{
		ID:           "cp-456",
		ProviderName: "test-provider",
		DisplayName:  "Test Provider",
		StandardType: "OAUTH2",
		Scopes: Scopes{
			DisplayLabel: "Test Scopes",
			Scopes: []ScopeChild{
				{
					ScopeName:   "read",
					Required:    true,
					Recommended: false,
				},
			},
		},
		UserinfoFields: map[string]interface{}{
			"sub": "user_id",
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, "providers-srv/custom/test-provider") {
			t.Errorf("Expected providers-srv/custom/test-provider endpoint, got %s", r.URL.Path)
		}

		response := CustomProviderResponse{
			Success: true,
			Status:  200,
			Data:    expectedProvider,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	customProvider := NewCustomProvider(config)

	result, err := customProvider.GetCustomProvider(context.Background(), "test-provider")

	if err != nil {
		t.Fatalf("GetCustomProvider failed: %v", err)
	}

	if !result.Success {
		t.Error("Expected success to be true")
	}

	if result.Data.ProviderName != expectedProvider.ProviderName {
		t.Errorf("Expected ProviderName %s, got %s", expectedProvider.ProviderName, result.Data.ProviderName)
	}
}

func TestCustomProvider_GetCustomProvider_Error(t *testing.T) {
	server := NewMockServer(http.StatusNotFound, `{"error": "provider not found"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	customProvider := NewCustomProvider(config)

	_, err := customProvider.GetCustomProvider(context.Background(), "nonexistent")

	if err == nil {
		t.Error("Expected error for not found provider, got nil")
	}
}

func TestCustomProvider_DeleteCustomProvider_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE method, got %s", r.Method)
		}

		// Verify lowercase conversion
		if !strings.Contains(r.URL.Path, "providers-srv/custom/test-provider") {
			t.Errorf("Expected providers-srv/custom/test-provider endpoint, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	customProvider := NewCustomProvider(config)

	err := customProvider.DeleteCustomProvider(context.Background(), "TEST-PROVIDER")

	if err != nil {
		t.Fatalf("DeleteCustomProvider failed: %v", err)
	}
}

func TestCustomProvider_DeleteCustomProvider_CaseConversion(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify that provider name is converted to lowercase
		if !strings.Contains(r.URL.Path, "mixed-case-provider") {
			t.Errorf("Expected lowercase 'mixed-case-provider' in URL path, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	customProvider := NewCustomProvider(config)

	err := customProvider.DeleteCustomProvider(context.Background(), "Mixed-Case-Provider")

	if err != nil {
		t.Fatalf("DeleteCustomProvider with case conversion failed: %v", err)
	}
}

func TestCustomProvider_DeleteCustomProvider_Error(t *testing.T) {
	server := NewMockServer(http.StatusNotFound, `{"error": "provider not found"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	customProvider := NewCustomProvider(config)

	err := customProvider.DeleteCustomProvider(context.Background(), "nonexistent")

	if err == nil {
		t.Error("Expected error for not found provider, got nil")
	}
}

func TestCustomProvider_GetAll_Success(t *testing.T) {
	expectedProviders := []CustomProviderModel{
		{
			ID:           "cp-1",
			ProviderName: "provider-1",
			DisplayName:  "Provider 1",
			StandardType: "OAUTH2",
			Pkce:         true,
		},
		{
			ID:           "cp-2",
			ProviderName: "provider-2",
			DisplayName:  "Provider 2",
			StandardType: "SAML",
			Pkce:         false,
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, "providers-srv/custom") {
			t.Errorf("Expected providers-srv/custom endpoint, got %s", r.URL.Path)
		}

		response := AllCustomProviderResponse{
			Success: true,
			Status:  200,
			Data:    expectedProviders,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	customProvider := NewCustomProvider(config)

	result, err := customProvider.GetAll(context.Background())

	if err != nil {
		t.Fatalf("GetAll failed: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("Expected 2 providers, got %d", len(result))
	}

	if result[0].ProviderName != "provider-1" {
		t.Errorf("Expected first provider name 'provider-1', got %s", result[0].ProviderName)
	}

	if result[1].ProviderName != "provider-2" {
		t.Errorf("Expected second provider name 'provider-2', got %s", result[1].ProviderName)
	}
}

func TestCustomProvider_GetAll_EmptyResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := AllCustomProviderResponse{
			Success: true,
			Status:  200,
			Data:    []CustomProviderModel{},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	customProvider := NewCustomProvider(config)

	result, err := customProvider.GetAll(context.Background())

	if err != nil {
		t.Fatalf("GetAll failed: %v", err)
	}

	if len(result) != 0 {
		t.Errorf("Expected 0 providers, got %d", len(result))
	}
}

func TestCustomProvider_GetAll_Error(t *testing.T) {
	server := NewMockServer(http.StatusInternalServerError, `{"error": "server error"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	customProvider := NewCustomProvider(config)

	_, err := customProvider.GetAll(context.Background())

	if err == nil {
		t.Error("Expected error for server error, got nil")
	}
}

func TestCustomProvider_CreateCustomProvider_WithComplexScopes(t *testing.T) {
	providerWithScopes := &CustomProviderModel{
		ProviderName: "complex-scopes-provider",
		Scopes: Scopes{
			DisplayLabel: "Complex Scopes",
			Scopes: []ScopeChild{
				{
					ScopeName:   "openid",
					Required:    true,
					Recommended: true,
				},
				{
					ScopeName:   "profile",
					Required:    false,
					Recommended: true,
				},
				{
					ScopeName:   "email",
					Required:    false,
					Recommended: false,
				},
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var receivedProvider CustomProviderModel
		json.Unmarshal(body, &receivedProvider)

		if len(receivedProvider.Scopes.Scopes) != 3 {
			t.Errorf("Expected 3 scopes, got %d", len(receivedProvider.Scopes.Scopes))
		}

		if receivedProvider.Scopes.DisplayLabel != "Complex Scopes" {
			t.Errorf("Expected DisplayLabel 'Complex Scopes', got %s", receivedProvider.Scopes.DisplayLabel)
		}

		response := CustomProviderResponse{
			Success: true,
			Status:  201,
			Data:    *providerWithScopes,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	customProvider := NewCustomProvider(config)

	result, err := customProvider.CreateCustomProvider(context.Background(), providerWithScopes)

	if err != nil {
		t.Fatalf("CreateCustomProvider with complex scopes failed: %v", err)
	}

	if len(result.Data.Scopes.Scopes) != 3 {
		t.Errorf("Expected 3 scopes in response, got %d", len(result.Data.Scopes.Scopes))
	}
}
