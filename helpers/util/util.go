package util

import (
	"io"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func InterfaceArray2StringArray(interfaceArray []interface{}) []string {
	result := make([]string, 0)
	//nolint:forcetypeassert
	for _, txt := range interfaceArray {
		if txt != nil {
			result = append(result, txt.(string))
		}
	}
	return result
}

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
	if values != nil {
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
