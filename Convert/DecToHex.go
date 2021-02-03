package Convert

import "strconv"

// DecToHex 将十进制转换为十六进制.
func (*Convert)DecToHex(number int64) string {
	return strconv.FormatInt(number, 16)
}
