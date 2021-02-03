package Convert

// StrToFloat64 将字符串转换为float64.其中"true", "TRUE", "True"为1.0 .
func (toolConvert *Convert)StrToFloat64(val string) (res float64) {
	if val == "true" || val == "TRUE" || val == "True" {
		res = 1.0
	} else {
		res = float64(toolConvert.StrToFloatStrict(val, 64, false))
	}

	return
}
