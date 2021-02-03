package Number

// MinInt 整数序列求最小值.
func (*Number)MinInt(nums ...int) (res int) {
	if len(nums) < 1 {
		return
	}
	res = nums[0]
	for _, v := range nums {
		if v < res {
			res = v
		}
	}

	return
}