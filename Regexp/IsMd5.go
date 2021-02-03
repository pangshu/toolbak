package Regexp

import "regexp"

// IsMd5 是否md5值.
func IsMd5(str string) bool {
	return str != "" && regexp.MustCompile(`^(?i)([0-9a-h]{32})$`).MatchString(str)
}
