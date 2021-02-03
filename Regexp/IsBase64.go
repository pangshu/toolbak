package Regexp

import "regexp"

// IsBase64 是否base64字符串.
func IsBase64(str string) bool {
	return str != "" && regexp.MustCompile("^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$").MatchString(str)
}
