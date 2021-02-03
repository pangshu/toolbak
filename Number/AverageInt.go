package Number

// AverageInt 对整数序列求平均值.
func (toolNumber *Number)AverageInt(nums ...int) (res float64) {
	length := len(nums)
	if length == 0 {
		return
	} else if length == 1 {
		res = float64(nums[0])
	} else {
		total := toolNumber.SumInt(nums...)
		res = float64(total) / float64(length)
	}

	return
}
