package Convert

// StrToInt32 将字符串转换为int32.
func (toolConvert *Convert)StrToInt32(val string) int32 {
	return int32(toolConvert.StrToIntStrict(val, 32, false))
}

