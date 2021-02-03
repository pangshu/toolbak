package Number

// Average 对整数序列求平均值.
func (toolNumber *Number)Average(nums ...interface{}) (res float64) {
	length := len(nums)
	if length == 0 {
		return
	} else if length == 1 {
		res,_ = toolNumber.ToFloat(nums[0])
	} else {
		var count int
		var err error
		var val, total float64
		for _, v := range nums {
			val,err = toolNumber.ToFloat(v)
			if err == nil {
				count++
				total += val
			}
		}

		res = total / float64(count)
	}

	return
}