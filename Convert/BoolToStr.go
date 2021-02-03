package Convert

// BoolToStr 将布尔值转换为字符串.
func (*Convert)BoolToStr(val bool) string {
	if val {
		return "true"
	}
	return "false"
}
