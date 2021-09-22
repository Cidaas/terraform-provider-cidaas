package cidaas

import (
	"context"
	"encoding/json"

	// "fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"terraform-provider-cidaas/helper_pkg/cidaas_sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceApp() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAppRead,
		Schema: map[string]*schema.Schema{
			"client_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"app_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"datatype": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"app_retrieval_status": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"app_retrieval_success": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func generateFlattenedAttributes(appData map[string]interface{}) (appAttributes []map[string]interface{}) {
	for key, value := range appData["data"].(map[string]interface{}) {
		attribute := make(map[string]interface{})
		attribute["name"] = key
		attribute["value"] = value
		attribute["datatype"] = reflect.TypeOf(value).Kind()
		stringKind := reflect.TypeOf("1").Kind()
		intKind := reflect.TypeOf(1).Kind()
		boolKind := reflect.TypeOf(true).Kind()
		if attribute["datatype"] == stringKind || attribute["datatype"] == intKind || attribute["datatype"] == boolKind {
			if attribute["datatype"] == intKind {
				attribute["value"] = strconv.Itoa(value.(int))
				attribute["datatype"] = "int"
			}
			if attribute["datatype"] == boolKind {
				attribute["value"] = strconv.FormatBool(value.(bool))
				attribute["datatype"] = "bool"
			}
			if attribute["datatype"] == stringKind {
				attribute["value"] = value.(string)
				attribute["datatype"] = "string"
			}
			appAttributes = append(appAttributes, attribute)
		}
	}

	return appAttributes

}

func dataSourceAppRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	cidaas_client := m.(cidaas_sdk.CidaasClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Warning Message Summary",
		Detail:   "This is the detailed warning message from providerConfigure",
	})

	client_id := d.Get("client_id").(string)

	// Pre Request setup
	url := "https://terraform-cidaas-test-free.cidaas.de/apps-srv/clients" + client_id

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	authAccessToken := cidaas_client.TokenData.AccessToken

	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Authorization", "Bearer "+authAccessToken)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Origin", "https://nightlybuild.cidaas.de")
	req.Header.Add("Referer", "https://nightlybuild.cidaas.de/apps-webapp/update")
	req.Header.Add("Accept-Language", "en,de;q=0.9,en-GB;q=0.8,en-US;q=0.7")
	req.Header.Add("Cookie",
		"cidaas_dr=b8cffdd0-2391-46da-8f2c-db92edd3371e;cidaas_sso=61a1babe-11fb-4282-9e0f-1c7cdc18c288; cidaas_sid=01c18815-f6cc-499c-acc3-5b3356ec34e3")

	r, err := client.Do(req)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while making request",
			Detail:   err.Error(),
		})
		return diags
	}
	defer r.Body.Close()

	appData := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&appData)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while decoding json into appData interface",
			Detail:   err.Error(),
		})
		return diags
	}

	appAttributes := generateFlattenedAttributes(appData)

	if err := d.Set("app_attributes", appAttributes); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting appData to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("app_retrieval_success", appData["success"]); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting app_retrieval_success to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("app_retrieval_status", appData["status"]); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Occured while setting app_retrieval_status to resourceData",
			Detail:   err.Error(),
		})
		return diags
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
