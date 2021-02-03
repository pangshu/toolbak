package Convert

// StrToUint64 将字符串转换为uint64.
func (toolConvert *Convert)StrToUint64(val string) uint64 {
	return uint64(toolConvert.StrToUintStrict(val, 64, false))
}
