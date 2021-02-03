package Number

// InRangeFloat32 数值是否在2个32位浮点数范围内.
func (*Number)InRangeFloat32(value, left, right float32) bool {
	if left > right {
		left, right = right, left
	}
	return value >= left && value <= right
}
