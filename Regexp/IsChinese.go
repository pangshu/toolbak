package Regexp

import "regexp"

// IsChinese 字符串是否全部中文.
func (*Regexp)IsChinese(str string) bool {
	return str != "" && regexp.MustCompile("^[\u4e00-\u9fa5]+$").MatchString(str)
}
