package String

// IsEmpty 字符串是否为空(包括空格).
func (toolbaktring *String)IsEmpty(str string) bool {
	if str == "" || len(toolbaktring.Trim(str)) == 0 {
		return true
	}

	return false
}
