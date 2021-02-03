package Os

import (
	"fmt"
	"strconv"
)

// ByteFormat 格式化文件比特大小.
// size为文件大小,decimal为要保留的小数位数,delimiter为数字和单位间的分隔符.
func (*Os)ByteFormat(size float64, decimal uint8, delimiter string) string {
	var arr = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB", "UnKnown"}
	var pos int = 0
	var j float64 = float64(size)
	for {
		if size >= 1024 {
			size = size / 1024
			j = j / 1024
			pos++
		} else {
			break
		}
	}
	if pos >= len(arr) { // fixed out index bug
		pos = len(arr) - 1
	}

	return fmt.Sprintf("%."+strconv.Itoa(int(decimal))+"f%s%s", j, delimiter, arr[pos])
}
