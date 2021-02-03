package Regexp

import "regexp"

// IsChineseName 字符串是否中文名称.
func IsChineseName(str string) bool {
	return str != "" && regexp.MustCompile("^[\u4e00-\u9fa5][.•·\u4e00-\u9fa5]{0,30}[\u4e00-\u9fa5]$").MatchString(str)
}

