package Number

// SumInt 整数求和.
func (*Number)SumInt(nums ...int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum
}
