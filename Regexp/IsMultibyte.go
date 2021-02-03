package Regexp

import "regexp"

// IsMultibyte 字符串是否含有多字节字符.
func IsMultibyte(str string) bool {
	return str != "" && regexp.MustCompile("[^\x00-\x7F]").MatchString(str)
}
