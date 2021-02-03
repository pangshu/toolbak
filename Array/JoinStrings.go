package Array

import "strings"

// JoinStrings 使用分隔符delimiter连接字符串数组.效率比Implode高.
func (*Array)JoinStrings(strs []string, delimiter string) (res string) {
	length := len(strs)
	if length == 0 {
		return
	}

	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
		if length--; length > 0 {
			sb.WriteString(delimiter)
		}
	}
	res = sb.String()

	return
}
