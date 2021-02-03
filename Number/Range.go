package Number

// Range 根据范围创建数组,包含指定的元素.
// start为起始元素值,end为末尾元素值.若start<end,返回升序的数组;若start>end,返回降序的数组.
func (toolNumber *Number)Range(start, end int) []int {
	res := make([]int, toolNumber.AbsInt64(int64(end-start))+1)
	for i := range res {
		if end > start {
			res[i] = start + i
		} else {
			res[i] = start - i
		}
	}
	return res
}
