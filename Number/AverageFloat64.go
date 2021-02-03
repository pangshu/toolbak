package Number

// AverageFloat64 对浮点数序列求平均值.
func (toolNumber *Number)AverageFloat64(nums ...float64) (res float64) {
	length := len(nums)
	if length == 0 {
		return
	} else if length == 1 {
		res = nums[0]
	} else {
		total := toolNumber.SumFloat64(nums...)
		res = total / float64(length)
	}

	return
}
