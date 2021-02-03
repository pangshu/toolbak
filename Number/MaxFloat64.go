package Number

import "math"

// MaxFloat64 64位浮点数序列求最大值.
func (*Number)MaxFloat64(nums ...float64) (res float64) {
	if len(nums) < 1 {
		return
	}

	res = nums[0]
	for _, v := range nums {
		res = math.Max(res, v)
	}

	return
}
