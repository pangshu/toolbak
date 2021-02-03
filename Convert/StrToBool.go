package Convert

import "strconv"

// StrToBool 将字符串转换为布尔值.
// 1, t, T, TRUE, true, True 等字符串为真.
// 0, f, F, FALSE, false, False 等字符串为假.
func (*Convert)StrToBool(val string) (res bool) {
	if val != "" {
		res, _ = strconv.ParseBool(val)
	}

	return
}
