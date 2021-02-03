package Convert

// StrToInt16 将字符串转换为int16.
func (toolConvert *Convert)StrToInt16(val string) int16 {
	return int16(toolConvert.StrToIntStrict(val, 16, false))
}
