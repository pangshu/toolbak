package Convert

import "strconv"

// OctToDec 将八进制转换为十进制.
func (*Convert)OctToDec(str string) (int64, error) {
	start := 0
	if len(str) > 1 && str[0:1] == "0" {
		start = 1
	}

	return strconv.ParseInt(str[start:], 8, 0)
}
