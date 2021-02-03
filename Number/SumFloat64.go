package Number

// SumFloat64 浮点数求和.
func (*Number)SumFloat64(nums ...float64) float64 {
	var sum float64
	for _, v := range nums {
		sum += v
	}
	return sum
}
