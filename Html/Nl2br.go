package Html

import "strings"

// Nl2br 将换行符转换为br标签.
func Nl2br(html string) string {
	if html == "" {
		return ""
	}
	return strings.Replace(html, "\n", "<br />", -1)
}
