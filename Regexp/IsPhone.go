package Regexp

import "regexp"

// IsPhone 是否电话号码(手机或固话).
func IsPhone(str string) bool {
	return str != "" && regexp.MustCompile(`(` + `^1[3-9]\d{9}$` + `)|(` + `^(010|02\d{1}|0[3-9]\d{2})-\d{7,9}(-\d+)?$` + `)`).MatchString(str)
}
