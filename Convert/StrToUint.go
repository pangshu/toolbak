package Convert

// StrToUint 将字符串转换为uint.
func (toolConvert *Convert)StrToUint(val string) uint {
	return uint(toolConvert.StrToUintStrict(val, 0, false))
}

