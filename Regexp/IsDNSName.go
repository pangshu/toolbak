package Regexp

import (
	"regexp"
	"strings"
	"github.com/pangshu/toolbak/Net"
)

// IsDNSName 是否DNS名称.
func IsDNSName(str string) bool {
	if str == "" || len(strings.Replace(str, ".", "", -1)) > 255 {
		// constraints already violated
		return false
	}
	return !Net.IsIP(str) && regexp.MustCompile(`^([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[\._]?$`).MatchString(str)
}

