package Convert

import "strconv"

// DecToBin 将十进制转换为二进制.
func (*Convert)DecToBin(number int64) string {
	return strconv.FormatInt(number, 2)
}
