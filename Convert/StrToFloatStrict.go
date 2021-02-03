package Convert

import "strconv"

// StrToFloatStrict 严格将字符串转换为浮点型.
// bitSize为类型位数,strict为是否严格检查.
func (*Convert)StrToFloatStrict(val string, bitSize int, strict bool) float64 {
	res, err := strconv.ParseFloat(val, bitSize)
	if err != nil {
		if strict {
			panic(err)
		}
	}
	return res
}
