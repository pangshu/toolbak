package Convert

import "strconv"

// DecToOct 将十进制转换为八进制.
func (*Convert)DecToOct(number int64) string {
	return strconv.FormatInt(number, 8)
}
