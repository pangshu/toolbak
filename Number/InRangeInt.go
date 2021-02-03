package Number

// InRangeInt 数值是否在2个整数范围内.
func (*Number)InRangeInt(value, left, right int) bool {
	if left > right {
		left, right = right, left
	}
	return value >= left && value <= right
}

