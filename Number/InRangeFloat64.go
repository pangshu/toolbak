package Number

// InRangeFloat64 数值是否在2个64位浮点数范围内.
func (*Number)InRangeFloat64(value, left, right float64) bool {
	if left > right {
		left, right = right, left
	}
	return value >= left && value <= right
}
