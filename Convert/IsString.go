package Convert

// IsString 变量是否字符串.
func (toolConvert *Convert)IsString(val interface{}) bool {
	return toolConvert.Gettype(val) == "string"
}
