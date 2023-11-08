package utils

import "encoding/json"

func ToJson(data interface{}) string {
	if data == nil {
		return ""
	}
	b, _ := json.Marshal(data)
	return string(b)
}
