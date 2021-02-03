package Convert

// StrToUint16 将字符串转换为uint16.
func (toolConvert *Convert)StrToUint16(val string) uint16 {
	return uint16(toolConvert.StrToUintStrict(val, 16, false))
}
