package Convert

import "strconv"

// StrToIntStrict 严格将字符串转换为有符号整型.
// bitSize为类型位数,strict为是否严格检查.
func (*Convert)StrToIntStrict(val string, bitSize int, strict bool) int64 {
	res, err := strconv.ParseInt(val, 0, bitSize)
	if err != nil {
		if strict {
			panic(err)
		}
	}
	return res
}
