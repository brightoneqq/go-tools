package gstring

import "encoding/json"

func IsEmpty(content string) bool {
	return content == ""
}

func IsNotEmpty(content string) bool {
	return content != ""
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func ToByteArray(content string) []byte {
	return []byte(content)
}
