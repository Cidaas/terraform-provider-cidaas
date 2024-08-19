package cidaas

import (
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
	ClientConfig
}
type HostedPageService interface {
	Upsert(hp HostedPageModel) (*HostedPageResponse, error)
	Get(hpGroupName string) (*HostedPageResponse, error)
	Delete(hpGroupName string) error
}

func NewHostedPage(clientConfig ClientConfig) HostedPageService {
	return &HostedPage{clientConfig}
}

func (hp *HostedPage) Upsert(hpm HostedPageModel) (*HostedPageResponse, error) {
	var response HostedPageResponse
	url := fmt.Sprintf("%s/%s", hp.BaseURL, "hostedpages-srv/hpgroup")
	httpClient := util.NewHTTPClient(url, http.MethodPost, hp.AccessToken)

	res, err := httpClient.MakeRequest(hpm)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (hp *HostedPage) Get(hpGroupName string) (*HostedPageResponse, error) {
	var response HostedPageResponse
	url := fmt.Sprintf("%s/%s/%s", hp.BaseURL, "hostedpages-srv/hpgroup", hpGroupName)
	httpClient := util.NewHTTPClient(url, http.MethodGet, hp.AccessToken)

	res, err := httpClient.MakeRequest(nil)
	if err = util.HandleResponseError(res, err); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = util.ProcessResponse(res, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (hp *HostedPage) Delete(hpGroupName string) error {
	url := fmt.Sprintf("%s/%s/%s", hp.BaseURL, "hostedpages-srv/hpgroup", hpGroupName)
	httpClient := util.NewHTTPClient(url, http.MethodDelete, hp.AccessToken)

	res, err := httpClient.MakeRequest(nil)
	if err = util.HandleResponseError(res, err); err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
