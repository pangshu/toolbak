package Convert

// StrToUint8 将字符串转换为uint8.
func (toolConvert *Convert)StrToUint8(val string) uint8 {
	return uint8(toolConvert.StrToUintStrict(val, 8, false))
}
