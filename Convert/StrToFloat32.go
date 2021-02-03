package Convert

// StrToFloat32 将字符串转换为float32.
func (toolConvert *Convert)StrToFloat32(val string) float32 {
	return float32(toolConvert.StrToFloatStrict(val, 32, false))
}
