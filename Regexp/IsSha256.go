package Regexp

import "regexp"

// IsSha256 是否Sha256值.
func IsSha256(str string) bool {
	return str != "" && regexp.MustCompile(`^(?i)([0-9a-h]{64})$`).MatchString(str)
	return str != "" && regexp.MustCompile(`^(?i)([0-9a-h]{128})$`).MatchString(str)
}
