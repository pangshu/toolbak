package Number

// AbsInt 整型取绝对值.
func (*Number)AbsInt64(number int64) int64 {
	r := number >> 63
	return (number ^ r) - r
}
