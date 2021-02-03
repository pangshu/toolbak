package Regexp

import "regexp"

// IsWhitespaces 是否全部空白字符,不包括空字符串.
func IsWhitespaces(str string) bool {
	return str != "" && regexp.MustCompile("^[[:space:]]+$").MatchString(str)
}
