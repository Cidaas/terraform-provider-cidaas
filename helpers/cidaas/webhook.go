package cidaas

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
)

var (
	AllowedAuthType          = []string{"APIKEY", "TOTP", "CIDAAS_OAUTH2"}
	AllowedKeyPlacementValue = []string{"query", "header"}
	AllowedEvents            = []string{"GROUP_USER_ROLE_REMOVED", "PROFILE_IMAGE_REMOVED", "GROUP_TYPE_DELETED", "PASSWORD_RESET_INITIATE", "SCOPE_CREATED",
		"CHECKOUT_SESSION_ASYNC_PAYMENT_SUCCEEDED", "LOGOUT", "LOGIN_WITH_SOCIAL", "ACCOUNT_DELETION_SCHEDULED", "APP_CREATED", "TEMPLATE_UPDATED",
		"EMAIL_CHANGED", "SCOPE_DELETED", "GEOFENCE_EXIT", "HOSTED_PAGE_DELETED", "CAPTCHA_CREATED", "CUSTOM_TEMPLATE_DELETED", "LOGIN_WITH_CIDAAS",
		"LOGIN_FAILURE", "DEVICE_DELETED", "USER_REGION_ENDED", "GROUP_TYPE_CREATED", "CUSTOM_CODE_VERIFICATION_TRIGGERED", "USER_REGION_IN_PROGRESS",
		"HOSTED_PAGE_CREATED", "REGISTRATION_FIELD_UPDATED", "USER_DEVICE_LINK_DELETED", "SCOPE_UPDATED", "FIELDSETUP_ADDED", "FIELDSETUP_MODIFIED",
		"ACCOUNT_CREATED_WITH_SOCIAL_IDENTITY", "CUSTOM_TEMPLATE_UPDATED", "IVR_TRIGGERED", "CONSENT_REJECTED", "ACCOUNT_MODIFIED", "PASSWORD_RESET",
		"PROFILE_IMAGE_UPDATED", "GROUP_TYPE_MODIFIED", "IDVALIDATOR_BTX_FINISHED", "HOSTED_PAGE_MODIFIED", "WEBHOOK_DELETED", "PUSH_SENT", "WEBHOOK_UPDATED",
		"CHECKOUT_SESSION_COMPLETED", "GROUP_ADMIN_ADDED", "DEVICE_CREATED", "MFA_REQUIRED", "ACCOUNT_DEACTIVATED", "INVALID_CLIENT_SECRET_REQUESTED",
		"REGISTRATION_FIELD_DELETED", "PHYSICAL_VERIFICATION_CONFIG", "ACCOUNT_CREATED_WITH_CIDAAS_IDENTITY", "ACCOUNT_MOBILE_NO_UNVERIFIED", "HOSTED_PAGE_UPDATED",
		"CAPTCHA_DELETED", "GROUP_CREATED", "ACCOUNT_EMAIL_VERIFIED", "ACCOUNT_ACTIVATED", "DEVICE_UPDATED", "INVITE_USER", "CAPTCHA_UPDATED", "PHYSICAL_VERIFICATION",
		"INVITE_ACCEPTED", "REGISTRATION_FIELD_CREATED", "GROUP_DELETED", "USER_DEVICE_LINK_CREATED", "COMBINED", "ACCOUNT_CIDAASIDENTITY_REMOVED", "CUSTOM_TEMPLATE_CREATED",
		"WEBHOOK_CREATED", "PASS_UPDATED", "GROUP_USER_ROLE_UPDATED", "ACCOUNT_CONFLICT", "ROLE_UPDATED", "SOCIAL_PROVIDER_ENABLED", "APP_DELETED", "CHECKOUT_SESSION_EXPIRED",
		"PASSWORD_CHANGED", "ACCOUNT_SOCIALIDENTITY_ADDED", "ACCOUNT_DELETED", "GROUP_NEW_USER_ADDED", "EMAIL_SENT", "ACCESS_TOKEN_OBTAINED", "SMS_SENT", "ROLE_DELETED",
		"GROUP_MODIFIED", "INVALID_REDIRECT_URI_REQUESTED", "ACCOUNT_SOCIALIDENTITY_REMOVED", "PHYSICAL_VERIFICATION_REMOVED", "PASS_CREATED", "PASS_DELETED",
		"IDVALIDATOR_CASE_STARTED", "GROUP_USER_ROLE_ADDED", "ACCOUNT_EMAIL_UNVERIFIED", "CONSENT_ACCEPTED", "FIELDSETUP_DELETED", "CHECKOUT_SESSION_ASYNC_PAYMENT_FAILED",
		"INVALID_CODE_VERIFIER_REQUESTED", "ACCOUNT_LOCKED", "NON_APPROVED_SCOPES_REQUESTED", "USER_REGION_STARTED", "ACCOUNT_MOBILE_NO_VERIFIED", "APP_MODIFIED",
		"ACCOUNT_CIDAASIDENTITY_ADDED", "ROLE_CREATED", "SOCIAL_PROVIDER_DISABLED", "IDVALIDATOR_CASE_FINISHED", "GEOFENCE_ENTER", "IDVALIDATOR_VALIDATION_FINISHED",
		"GROUP_FIRST_ADMIN_ADDED", "GROUP_USER_REMOVED", "IDVALIDATOR_DOCSIGN_FINISHED",
	}
)

type WebhookModel struct {
	ID                string        `json:"_id,omitempty"`
	AuthType          string        `json:"auth_type,omitempty"`
	URL               string        `json:"url,omitempty"`
	Events            []string      `json:"events,omitempty"`
	APIKeyDetails     APIKeyDetails `json:"apikeyDetails,omitempty"`
	TotpDetails       TotpDetails   `json:"totpDetails,omitempty"`
	CidaasAuthDetails AuthDetails   `json:"cidaasAuthDetails,omitempty"`
	Disable           bool          `json:"disable"`
	CreatedTime       string        `json:"createdTime,omitempty"`
	UpdatedTime       string        `json:"updatedTime,omitempty"`
}

type APIKeyDetails struct {
	ApikeyPlaceholder string `json:"apikey_placeholder,omitempty"`
	ApikeyPlacement   string `json:"apikey_placement,omitempty"`
	Apikey            string `json:"apikey,omitempty"`
}

type TotpDetails struct {
	TotpPlaceholder string `json:"totp_placeholder,omitempty"`
	TotpPlacement   string `json:"totp_placement,omitempty"`
	TotpKey         string `json:"totpkey,omitempty"`
}
type AuthDetails struct {
	ClientID string `json:"client_id,omitempty"`
}

type WebhookResponse struct {
	Success bool         `json:"success,omitempty"`
	Status  int          `json:"status,omitempty"`
	Data    WebhookModel `json:"data,omitempty"`
}

var _ WebhookService = &Webhook{}

type Webhook struct {
	HTTPClient util.HTTPClientInterface
}
type WebhookService interface {
	Upsert(wb WebhookModel) (*WebhookResponse, error)
	Get(id string) (*WebhookResponse, error)
	Delete(id string) error
}

func NewWebhook(httpClient util.HTTPClientInterface) WebhookService {
	return &Webhook{HTTPClient: httpClient}
}

func (w *Webhook) Upsert(wb WebhookModel) (*WebhookResponse, error) {
	w.HTTPClient.SetURL(fmt.Sprintf("%s/%s", w.HTTPClient.GetHost(), "webhook-srv/webhook"))
	w.HTTPClient.SetMethod(http.MethodPost)
	res, err := w.HTTPClient.MakeRequest(wb)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response WebhookResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (w *Webhook) Get(id string) (*WebhookResponse, error) {
	w.HTTPClient.SetURL(fmt.Sprintf("%s/%s?id=%s", w.HTTPClient.GetHost(), "webhook-srv/webhook", id))
	w.HTTPClient.SetMethod(http.MethodGet)
	res, err := w.HTTPClient.MakeRequest(nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response WebhookResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (w *Webhook) Delete(id string) error {
	w.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s", w.HTTPClient.GetHost(), "webhook-srv/webhook", id))
	w.HTTPClient.SetMethod(http.MethodDelete)
	res, err := w.HTTPClient.MakeRequest(nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
