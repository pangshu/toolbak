package Regexp

import "regexp"

// IsSha512 是否Sha512值.
func IsSha512(str string) bool {
	return str != "" && regexp.MustCompile(`^(?i)([0-9a-h]{128})$`).MatchString(str)
}
