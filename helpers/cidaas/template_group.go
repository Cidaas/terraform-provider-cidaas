package cidaas

import (
	"encoding/json"
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
	HTTPClient util.HTTPClientInterface
}

type TemplateGroupService interface {
	Create(tg TemplateGroupModel) (*TemplateGroupResponse, error)
	Update(tg TemplateGroupModel) (*TemplateGroupResponse, error)
	Get(groupID string) (*TemplateGroupResponse, error)
	Delete(groupID string) error
}

func NewTemplateGroup(httpClient util.HTTPClientInterface) TemplateGroupService {
	return &TemplateGroup{HTTPClient: httpClient}
}

func (rf *TemplateGroup) Create(tg TemplateGroupModel) (*TemplateGroupResponse, error) {
	rf.HTTPClient.SetURL(fmt.Sprintf("%s/%s", rf.HTTPClient.GetHost(), "templates-srv/groups"))
	rf.HTTPClient.SetMethod(http.MethodPost)
	res, err := rf.HTTPClient.MakeRequest(tg)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response TemplateGroupResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (rf *TemplateGroup) Update(tg TemplateGroupModel) (*TemplateGroupResponse, error) {
	rf.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s", rf.HTTPClient.GetHost(), "templates-srv/groups", tg.GroupID))
	rf.HTTPClient.SetMethod(http.MethodPut)
	res, err := rf.HTTPClient.MakeRequest(tg)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response TemplateGroupResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (r *TemplateGroup) Get(groupID string) (*TemplateGroupResponse, error) {
	r.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s", r.HTTPClient.GetHost(), "templates-srv/groups", groupID))
	r.HTTPClient.SetMethod(http.MethodGet)
	res, err := r.HTTPClient.MakeRequest(nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response TemplateGroupResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (r *TemplateGroup) Delete(groupID string) error {
	r.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s", r.HTTPClient.GetHost(), "templates-srv/groups", groupID))
	r.HTTPClient.SetMethod(http.MethodDelete)
	res, err := r.HTTPClient.MakeRequest(nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
