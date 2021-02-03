package Array

import "sort"

// UniqueInts 移除整数数组中的重复值.
func (*Array)UniqueInts(ints []int) (res []int) {
	sort.Ints(ints)
	var last int
	for i, num := range ints {
		if i == 0 || num != last {
			res = append(res, num)
		}
		last = num
	}
	return
}
