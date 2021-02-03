package Regexp

import (
	"regexp"
	"strings"
)

// IsBase64Image 是否base64编码的图片.
func IsBase64Image(str string) bool {
	if str == "" || !strings.ContainsRune(str, ',') {
		return false
	}

	dataURI := strings.Split(str, ",")
	return regexp.MustCompile(`^data:\s*(image|img)\/(\w+);base64`).MatchString(dataURI[0]) && regexp.MustCompile("^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$").MatchString(dataURI[1])
}
