package Convert

// IsHex 是否十六进制字符串.
func (toolConvert *Convert)IsHex(str string) bool {
	_, err := toolConvert.HexToDec(str)
	return err == nil
}
