package Array

// InInt64Slice 是否在64位整型数组/切片内.
func (*Array)InInt64Slice(i int64, list []int64) bool {
	for _, item := range list {
		if item == i {
			return true
		}
	}
	return false
}
