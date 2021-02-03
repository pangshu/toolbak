package Regexp

import "regexp"

// IsRgbColor 检查字符串是否RGB颜色格式.
func IsRgbColor(str string) bool {
	return str != "" && regexp.MustCompile("^rgb\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*\\)$").MatchString(str)
}
