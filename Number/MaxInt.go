package Number

// MaxInt 整数序列求最大值.
func (*Number)MaxInt(nums ...int) (res int) {
	if len(nums) < 1 {
		return
	}

	res = nums[0]
	for _, v := range nums {
		if v > res {
			res = v
		}
	}

	return
}
