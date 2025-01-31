package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func responseToStringConvert(resp *http.Response) string {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	return string(bodyBytes)
}

func StringInSlice(s string, list []string) bool {
	for _, v := range list {
		if s == v {
			return true
		}
	}
	return false
}

func Int64ValueOrNull(value *int64) types.Int64 {
	if value != nil {
		return types.Int64Value(*value)
	}
	return types.Int64Null()
}

func StringValueOrNull(value *string) types.String {
	if value != nil {
		if *value == "" {
			return types.StringNull()
		}
		return types.StringValue(*value)
	}
	return types.StringNull()
}

func SetValueOrNull(values []string) types.Set {
	if len(values) > 0 {
		var temp []attr.Value
		for _, v := range values {
			temp = append(temp, types.StringValue(v))
		}
		return types.SetValueMust(types.StringType, temp)
	}
	return types.SetNull(types.StringType)
}

func BoolValueOrNull(value *bool) types.Bool {
	if value != nil {
		return types.BoolValue(*value)
	}
	return types.BoolNull()
}

func TimeValueOrNull(value *time.Time) types.String {
	if value != nil && !value.IsZero() {
		return types.StringValue(value.Format("2006-01-02T15:04:05Z"))
	}
	return types.StringNull()
}

func MapValueOrNull(value *map[string]string) (types.Map, diag.Diagnostics) {
	if value == nil || len(*value) < 1 {
		return types.MapNull(types.StringType), nil
	}
	mapAttributes := map[string]attr.Value{}
	for key, value := range *value {
		mapAttributes[key] = types.StringValue(value)
	}
	cf, diag := types.MapValue(types.StringType, mapAttributes)
	return cf, diag
}

func ProcessResponse(res *http.Response, target interface{}) error {
	if target != nil {
		err := json.NewDecoder(res.Body).Decode(target)
		if err != nil {
			return fmt.Errorf("failed to decode response body, %w", err)
		}
	}
	return nil
}

func HandleResponseError(res *http.Response, err error) error {
	if err != nil {
		if res != nil && res.Body != nil {
			defer res.Body.Close()
		}
		return err
	}
	return nil
}

func FormatErrorMessage(err error) string {
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	return ""
}

func SetValueOrEmpty(values []string) types.Set {
	if len(values) == 0 {
		return types.SetValueMust(types.StringType, []attr.Value{})
	}
	elements := make([]attr.Value, 0, len(values))
	for _, v := range values {
		elements = append(elements, types.StringValue(v))
	}
	return types.SetValueMust(types.StringType, elements)
}
