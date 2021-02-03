package Convert

// IsByte 变量是否字节切片.
func (toolConvert *Convert)IsByte(val interface{}) bool {
	return toolConvert.Gettype(val) == "[]uint8"
}

