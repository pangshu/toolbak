package Convert

// IsBool 是否布尔值.
func (*Convert)IsBool(val interface{}) bool {
	return val == true || val == false
}
