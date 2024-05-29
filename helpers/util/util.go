package util

import (
	"io"
	"net/http"
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
