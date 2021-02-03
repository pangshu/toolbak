package Convert

// IsBinary 字符串是否二进制.
func (*Convert)IsBinary(s string) bool {
	for _, b := range s {
		if 0 == b {
			return true
		}
	}
	return false
}
