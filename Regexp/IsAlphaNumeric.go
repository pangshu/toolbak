package Regexp

import "regexp"

// IsAlphaNumeric 是否字母或数字.
func IsAlphaNumeric(str string) bool {
	//regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(str)
	return str != "" && regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(str)
}
