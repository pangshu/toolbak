package Array

import (
	"strconv"
	"strings"
)

// JoinInts 使用分隔符delimiter连接整数数组.
func (*Array)JoinInts(ints []int, delimiter string) (res string) {
	length := len(ints)
	if length == 0 {
		return
	}

	var sb strings.Builder
	for _, num := range ints {
		sb.WriteString(strconv.Itoa(num))
		if length--; length > 0 {
			sb.WriteString(delimiter)
		}
	}
	res = sb.String()

	return
}
