package cidaas

import (
	"fmt"
	"net/http"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
)

type TemplateGroupResponse struct {
	Success bool               `json:"success"`
	Status  int                `json:"status"`
	Data    TemplateGroupModel `json:"data"`
}

type TemplateGroupModel struct {
	ID                string             `json:"id,omitempty"`
	GroupID           string             `json:"group_id,omitempty"`
	SenderConfig      *SenderConfig      `json:"sender_config,omitempty"`
	EmailSenderConfig *EmailSenderConfig `json:"email_sender_config,omitempty"`
	SMSSenderConfig   *SMSSenderConfig   `json:"sms_sender_config,omitempty"`
	IVRSenderConfig   *IVRSenderConfig   `json:"ivr_sender_config,omitempty"`
	PushSenderConfig  *IVRSenderConfig   `json:"push_sender_config,omitempty"`
}

type SenderConfig struct {
	ID        string `json:"id,omitempty"`
	FromEmail string `json:"from_email,omitempty"`
	FromName  string `json:"from_name,omitempty"`
}

type EmailSenderConfig struct {
	ID          string   `json:"id,omitempty"`
	FromEmail   string   `json:"from_email,omitempty"`
	FromName    string   `json:"from_name,omitempty"`
	ReplyTo     string   `json:"reply_to,omitempty"`
	SenderNames []string `json:"sender_names,omitempty"`
}

type SMSSenderConfig struct {
	ID          string   `json:"id,omitempty"`
	FromName    string   `json:"from_name,omitempty"`
	SenderNames []string `json:"sender_names,omitempty"`
}

type IVRSenderConfig struct {
	ID          string   `json:"id,omitempty"`
	SenderNames []string `json:"sender_names,omitempty"`
}

var _ TemplateGroupService = &TemplateGroup{}

type TemplateGroup struct {
	ClientConfig
}

type TemplateGroupService interface {
	Create(tg TemplateGroupModel) (*TemplateGroupResponse, error)
	Update(tg TemplateGroupModel) (*TemplateGroupResponse, error)
	Get(groupID string) (*TemplateGroupResponse, error)
	Delete(groupID string) error
}

func NewTemplateGroup(clientConfig ClientConfig) TemplateGroupService {
	return &TemplateGroup{clientConfig}
}

func (t *TemplateGroup) Create(tg TemplateGroupModel) (*TemplateGroupResponse, error) {
	var response TemplateGroupResponse
	url := fmt.Sprintf("%s/%s", t.BaseURL, "templates-srv/groups")
	httpClient := util.NewHTTPClient(url, http.MethodPost, t.AccessToken)

	res, err := httpClient.MakeRequest(tg)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (t *TemplateGroup) Update(tg TemplateGroupModel) (*TemplateGroupResponse, error) {
	var response TemplateGroupResponse
	url := fmt.Sprintf("%s/%s/%s", t.BaseURL, "templates-srv/groups", tg.GroupID)
	httpClient := util.NewHTTPClient(url, http.MethodPut, t.AccessToken)

	res, err := httpClient.MakeRequest(tg)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (t *TemplateGroup) Get(groupID string) (*TemplateGroupResponse, error) {
	var response TemplateGroupResponse
	url := fmt.Sprintf("%s/%s/%s", t.BaseURL, "templates-srv/groups", groupID)
	httpClient := util.NewHTTPClient(url, http.MethodGet, t.AccessToken)

	res, err := httpClient.MakeRequest(nil)
	if res.StatusCode == http.StatusNoContent {
		resp := &TemplateGroupResponse{
			Status: http.StatusNoContent,
		}
		return resp, fmt.Errorf("template group not found by the provider group_id  %s", groupID)
	}
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (t *TemplateGroup) Delete(groupID string) error {
	url := fmt.Sprintf("%s/%s/%s", t.BaseURL, "templates-srv/groups", groupID)
	httpClient := util.NewHTTPClient(url, http.MethodDelete, t.AccessToken)

	res, err := httpClient.MakeRequest(nil)
	if err = util.HandleResponseError(res, err); err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
