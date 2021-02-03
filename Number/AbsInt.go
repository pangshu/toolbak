package Number

// AbsInt 整型取绝对值.
func (*Number)AbsInt(number int) int {
	r := number >> 63
	return (number ^ r) - r
}
