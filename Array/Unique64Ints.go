package Array

// Unique64Ints 移除64位整数数组中的重复值.
func (*Array)Unique64Ints(ints []int64) (res []int64) {
	seen := make(map[int64]bool)
	for _, num := range ints {
		if _, ok := seen[num]; !ok {
			seen[num] = true
			res = append(res, num)
		}
	}
	return
}
