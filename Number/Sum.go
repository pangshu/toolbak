package Number

// Sum 对任意类型序列中的数值类型求和,忽略非数值的.
func (toolNumber *Number)Sum(nums ...interface{}) (res float64) {
	var val float64
	for _, v := range nums {
		val,_ = toolNumber.ToFloat(v)
		if val > 0 {
			res += val
		}
	}

	return
}
