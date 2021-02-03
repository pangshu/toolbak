package Regexp

import (
	"regexp"
)

// AtWho 查找被@的用户名.minLen为用户名最小长度,默认5.
func AtWho(content string) []string {
	reg := regexp.MustCompile("@([\\w\\x{2e80}-\\x{9fff}\\-]+)").FindAllString(content, -1)
	return reg
}

//
//func AtWho(text string, minLen ...int) []string {
//	var result = []string{}
//	var username string
//	var min int = 5
//	if len(minLen) > 0 && minLen[0] > 0 {
//		min = minLen[0]
//	}
//
//	for _, line := range strings.Split(text, "\n") {
//		if len(line) == 0 {
//			continue
//		}
//		for {
//			index := strings.Index(line, "@")
//			if index == -1 {
//				break
//			} else if index > 0 {
//				r := rune(line[index-1])
//				if unicode.IsUpper(r) || unicode.IsLower(r) {
//					line = line[index+1:]
//				} else {
//					line = line[index:]
//				}
//			} else if index == 0 {
//				// the "@" is first characters
//				endIndex := strings.Index(line, " ")
//				if endIndex == -1 {
//					username = line[1:]
//				} else {
//					username = line[1:endIndex]
//				}
//
//				if len(username) >= min && regexp.MustCompile(`^[a-zA-Z0-9_.]+$`).MatchString(username) && !Array.InStringSlice(username, result) {
//					result = append(result, username)
//				}
//
//				if endIndex == -1 {
//					break
//				}
//
//				line = line[endIndex:]
//			}
//		}
//	}
//
//	return result
//}
