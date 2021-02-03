package Number

import "math"

// Min 取出任意类型中数值类型的最小值,无数值类型则为0.
func (toolNumber *Number)Min(nums ...interface{}) (res float64) {
	if len(nums) < 1 {
		return
	}

	var err error
	var val float64
	res, _ = toolNumber.ToFloat(nums[0])
	for _, v := range nums {
		val, err = toolNumber.ToFloat(v)
		if err == nil {
			res = math.Min(res, val)
		}
	}

	return
}