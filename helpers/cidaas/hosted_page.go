package cidaas

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Cidaas/terraform-provider-cidaas/helpers/util"
)

type HostedPageModel struct {
	ID            string           `json:"_id,omitempty"`
	GroupOwner    string           `json:"groupOwner,omitempty"`
	DefaultLocale string           `json:"default_locale,omitempty"`
	HostedPages   []HostedPageData `json:"hosted_pages,omitempty"`
	CreatedTime   string           `json:"createdTime,omitempty"`
	UpdatedTime   string           `json:"updatedTime,omitempty"`
}

type HostedPageData struct {
	HostedPageID string `json:"hosted_page_id,omitempty"`
	Content      string `json:"content,omitempty"`
	Locale       string `json:"locale,omitempty"`
	URL          string `json:"url,omitempty"`
}

type HostedPageResponse struct {
	Success bool            `json:"success,omitempty"`
	Status  int             `json:"status,omitempty"`
	Data    HostedPageModel `json:"data,omitempty"`
}

type HostedPage struct {
	HTTPClient util.HTTPClientInterface
}
type HostedPageService interface {
	Upsert(hp HostedPageModel) (*HostedPageResponse, error)
	Get(hpGroupName string) (*HostedPageResponse, error)
	Delete(hpGroupName string) error
}

func NewHostedPage(httpClient util.HTTPClientInterface) HostedPageService {
	return &HostedPage{HTTPClient: httpClient}
}

func (hp *HostedPage) Upsert(hpm HostedPageModel) (*HostedPageResponse, error) {
	hp.HTTPClient.SetURL(fmt.Sprintf("%s/%s", hp.HTTPClient.GetHost(), "hostedpages-srv/hpgroup"))
	hp.HTTPClient.SetMethod(http.MethodPost)
	res, err := hp.HTTPClient.MakeRequest(hpm)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response HostedPageResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (hp *HostedPage) Get(hpGroupName string) (*HostedPageResponse, error) {
	hp.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s", hp.HTTPClient.GetHost(), "hostedpages-srv/hpgroup", hpGroupName))
	hp.HTTPClient.SetMethod(http.MethodGet)
	res, err := hp.HTTPClient.MakeRequest(nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response HostedPageResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body, %w", err)
	}
	return &response, nil
}

func (hp *HostedPage) Delete(hpGroupName string) error {
	hp.HTTPClient.SetURL(fmt.Sprintf("%s/%s/%s", hp.HTTPClient.GetHost(), "hostedpages-srv/hpgroup", hpGroupName))
	hp.HTTPClient.SetMethod(http.MethodDelete)
	res, err := hp.HTTPClient.MakeRequest(nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
