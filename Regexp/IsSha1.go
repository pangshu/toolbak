package Regexp

import "regexp"

// IsSha1 是否Sha1值.
func IsSha1(str string) bool {
	return str != "" && regexp.MustCompile(`^(?i)([0-9a-h]{40})$`).MatchString(str)
}

