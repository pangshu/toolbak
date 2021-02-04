package Regexp

import (
	"net"
	"github.com/pangshu/toolbak/Net"
)

// IsDialString 是否网络拨号字符串(形如127.0.0.1:80),用于net.Dial()检查.
func IsDialString(str string) bool {
	h, p, err := net.SplitHostPort(str)
	if err == nil && h != "" && p != "" && (IsDNSName(h) || Net.IsIP(h)) && Net.IsPort(p) {
		return true
	}

	return false
}
