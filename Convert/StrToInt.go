package Convert

import "strconv"

// StrToInt 将字符串转换为int.其中"true", "TRUE", "True"为1.
func (*Convert)StrToInt(val string) (res int) {
	if val == "true" || val == "TRUE" || val == "True" {
		res = 1
		return
	}

	res, _ = strconv.Atoi(val)
	return
}
