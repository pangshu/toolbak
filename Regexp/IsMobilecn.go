package Regexp

import "regexp"

// IsMobilecn 检查字符串是否中国大陆手机号.
func IsMobilecn(str string) bool {
	return str != "" && regexp.MustCompile(`^1[3-9]\d{9}$`).MatchString(str)
}

