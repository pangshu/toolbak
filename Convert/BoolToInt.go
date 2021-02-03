package Convert

// BoolToInt 将布尔值转换为整型.
func (*Convert)BoolToInt(val bool) int {
	if val {
		return 1
	}
	return 0
}
