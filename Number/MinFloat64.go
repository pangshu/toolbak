package Number

import "math"

// MinFloat64 64位浮点数序列求最小值.
func (*Number)MinFloat64(nums ...float64) (res float64) {
	if len(nums) < 1 {
		return
	}
	res = nums[0]
	for _, v := range nums {
		res = math.Min(res, v)
	}

	return
}
