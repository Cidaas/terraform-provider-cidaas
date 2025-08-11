// helpers/cidaas/hosted_page_test.go
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

func TestNewHostedPage(t *testing.T) {
	config := NewTestClientConfig("http://test.com")

	hostedPage := NewHostedPage(config)

	if hostedPage == nil {
		t.Fatal("Expected hosted page instance, got nil")
	}

	if hostedPage.BaseURL != config.BaseURL {
		t.Errorf("Expected BaseURL %s, got %s", config.BaseURL, hostedPage.BaseURL)
	}

	if hostedPage.AccessToken != config.AccessToken {
		t.Errorf("Expected AccessToken %s, got %s", config.AccessToken, hostedPage.AccessToken)
	}
}

func TestHostedPage_Upsert_Success(t *testing.T) {
	expectedHostedPage := HostedPageModel{
		ID:            "hp-123",
		GroupOwner:    "admin",
		DefaultLocale: "en-US",
		HostedPages: []HostedPageData{
			{
				HostedPageID: "login-page",
				Content:      "<html><body>Login Page</body></html>",
				Locale:       "en-US",
				URL:          "https://example.com/login",
			},
			{
				HostedPageID: "register-page",
				Content:      "<html><body>Register Page</body></html>",
				Locale:       "en-US",
				URL:          "https://example.com/register",
			},
		},
		CreatedTime: "2024-01-15T10:30:00Z",
		UpdatedTime: "2024-01-15T11:45:00Z",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, "hostedpages-srv/hpgroup") {
			t.Errorf("Expected hostedpages-srv/hpgroup endpoint, got %s", r.URL.Path)
		}

		body, _ := io.ReadAll(r.Body)
		var receivedHostedPage HostedPageModel
		json.Unmarshal(body, &receivedHostedPage)

		if receivedHostedPage.GroupOwner != expectedHostedPage.GroupOwner {
			t.Errorf("Expected GroupOwner %s, got %s", expectedHostedPage.GroupOwner, receivedHostedPage.GroupOwner)
		}

		if len(receivedHostedPage.HostedPages) != 2 {
			t.Errorf("Expected 2 hosted pages, got %d", len(receivedHostedPage.HostedPages))
		}

		response := HostedPageResponse{
			Success: true,
			Status:  200,
			Data:    expectedHostedPage,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	hostedPage := NewHostedPage(config)

	result, err := hostedPage.Upsert(context.Background(), expectedHostedPage)

	if err != nil {
		t.Fatalf("Upsert failed: %v", err)
	}

	if !result.Success {
		t.Error("Expected success to be true")
	}

	if result.Data.GroupOwner != expectedHostedPage.GroupOwner {
		t.Errorf("Expected GroupOwner %s, got %s", expectedHostedPage.GroupOwner, result.Data.GroupOwner)
	}

	if len(result.Data.HostedPages) != 2 {
		t.Errorf("Expected 2 hosted pages in response, got %d", len(result.Data.HostedPages))
	}
}

func TestHostedPage_Upsert_Error(t *testing.T) {
	server := NewMockServer(http.StatusBadRequest, `{"error": "invalid hosted page"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	hostedPage := NewHostedPage(config)

	hostedPageModel := HostedPageModel{GroupOwner: "test"}
	_, err := hostedPage.Upsert(context.Background(), hostedPageModel)

	if err == nil {
		t.Error("Expected error for bad request, got nil")
	}
}

func TestHostedPage_Get_Success(t *testing.T) {
	expectedHostedPage := HostedPageModel{
		ID:            "hp-456",
		GroupOwner:    "user",
		DefaultLocale: "de-DE",
		HostedPages: []HostedPageData{
			{
				HostedPageID: "forgot-password",
				Content:      "<html><body>Forgot Password</body></html>",
				Locale:       "de-DE",
				URL:          "https://example.com/forgot-password",
			},
		},
		CreatedTime: "2024-01-10T09:00:00Z",
		UpdatedTime: "2024-01-10T09:30:00Z",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, "hostedpages-srv/hpgroup/test-group") {
			t.Errorf("Expected hostedpages-srv/hpgroup/test-group endpoint, got %s", r.URL.Path)
		}

		response := HostedPageResponse{
			Success: true,
			Status:  200,
			Data:    expectedHostedPage,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	hostedPage := NewHostedPage(config)

	result, err := hostedPage.Get(context.Background(), "test-group")

	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}

	if !result.Success {
		t.Error("Expected success to be true")
	}

	if result.Data.ID != expectedHostedPage.ID {
		t.Errorf("Expected ID %s, got %s", expectedHostedPage.ID, result.Data.ID)
	}

	if result.Data.DefaultLocale != expectedHostedPage.DefaultLocale {
		t.Errorf("Expected DefaultLocale %s, got %s", expectedHostedPage.DefaultLocale, result.Data.DefaultLocale)
	}
}

func TestHostedPage_Get_Error(t *testing.T) {
	server := NewMockServer(http.StatusNotFound, `{"error": "hosted page group not found"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	hostedPage := NewHostedPage(config)

	_, err := hostedPage.Get(context.Background(), "nonexistent-group")

	if err == nil {
		t.Error("Expected error for not found group, got nil")
	}
}

func TestHostedPage_Delete_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, "hostedpages-srv/hpgroup/test-group") {
			t.Errorf("Expected hostedpages-srv/hpgroup/test-group endpoint, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	hostedPage := NewHostedPage(config)

	err := hostedPage.Delete(context.Background(), "test-group")

	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
}

func TestHostedPage_Delete_Error(t *testing.T) {
	server := NewMockServer(http.StatusInternalServerError, `{"error": "server error"}`)
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	hostedPage := NewHostedPage(config)

	err := hostedPage.Delete(context.Background(), "test-group")

	if err == nil {
		t.Error("Expected error for server error, got nil")
	}
}

func TestHostedPage_Upsert_MinimalData(t *testing.T) {
	minimalHostedPage := HostedPageModel{
		GroupOwner:    "minimal",
		DefaultLocale: "en",
		HostedPages:   []HostedPageData{},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var receivedHostedPage HostedPageModel
		json.Unmarshal(body, &receivedHostedPage)

		if receivedHostedPage.GroupOwner != minimalHostedPage.GroupOwner {
			t.Errorf("Expected GroupOwner %s, got %s", minimalHostedPage.GroupOwner, receivedHostedPage.GroupOwner)
		}

		if len(receivedHostedPage.HostedPages) != 0 {
			t.Errorf("Expected 0 hosted pages, got %d", len(receivedHostedPage.HostedPages))
		}

		response := HostedPageResponse{
			Success: true,
			Status:  200,
			Data:    minimalHostedPage,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := NewTestClientConfig(server.URL)
	hostedPage := NewHostedPage(config)

	result, err := hostedPage.Upsert(context.Background(), minimalHostedPage)

	if err != nil {
		t.Fatalf("Upsert with minimal data failed: %v", err)
	}

	if len(result.Data.HostedPages) != 0 {
		t.Errorf("Expected 0 hosted pages in response, got %d", len(result.Data.HostedPages))
	}
}
