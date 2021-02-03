package Convert

// StrToUint32 将字符串转换为uint32.
func (toolConvert *Convert)StrToUint32(val string) uint32 {
	return uint32(toolConvert.StrToUintStrict(val, 32, false))
}
