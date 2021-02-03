package Regexp

import (
	"regexp"
	"strings"
)

// IsHexcolor 检查是否十六进制颜色,并返回带"#"的修正值.
func IsHexcolor(str string) (bool, string) {
	chk := str != "" && regexp.MustCompile(`^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$`).MatchString(str)
	if chk && !strings.ContainsRune(str, '#') {
		str = "#" + strings.ToUpper(str)
	}
	return chk, str
}
