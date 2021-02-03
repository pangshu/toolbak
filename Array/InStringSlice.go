package Array

// InStringSlice 是否在字符串数组/切片内.
func (*Array)InStringSlice(str string, list []string) bool {
	for _, item := range list {
		if item == str {
			return true
		}
	}
	return false
}
