package Convert

import "strconv"

// StrToUintStrict 严格将字符串转换为无符号整型,bitSize为类型位数,strict为是否严格检查
func (*Convert)StrToUintStrict(val string, bitSize int, strict bool) uint64 {
	res, err := strconv.ParseUint(val, 0, bitSize)
	if err != nil {
		if strict {
			panic(err)
		}
	}
	return res
}
