package Convert

// StrToInt64 将字符串转换为int64.
func (toolConvert *Convert)StrToInt64(val string) int64 {
	return toolConvert.StrToIntStrict(val, 64, false)
}
