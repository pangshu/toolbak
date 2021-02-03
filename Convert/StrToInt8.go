package Convert

// StrToInt8 将字符串转换为int8.
func (toolConvert *Convert)StrToInt8(val string) int8 {
	return int8(toolConvert.StrToIntStrict(val, 8, false))
}

