package Array

// InIntSlice 是否在整型数组/切片内.
func (*Array)InIntSlice(i int, list []int) bool {
	for _, item := range list {
		if item == i {
			return true
		}
	}
	return false
}
