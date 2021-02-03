package Regexp

import "regexp"

// IsWord 是否词语(不以下划线开头的中文、英文、数字、下划线).
func IsWord(str string) bool {
	return str != "" && regexp.MustCompile("^[a-zA-Z0-9\u4e00-\u9fa5][a-zA-Z0-9_\u4e00-\u9fa5]+$").MatchString(str)
}
