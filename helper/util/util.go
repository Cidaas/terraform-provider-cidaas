package util

import (
	"fmt"
	"io"
	"net/http"
)

func InterfaceArray2StringArray(interfaceArray []interface{}) []string {
	result := make([]string, 0)
	for _, txt := range interfaceArray {
		result = append(result, txt.(string))
	}
	return result
}

func responseToStringConvert(resp *http.Response) string {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf(err.Error())
	}
	return string(bodyBytes)
}
