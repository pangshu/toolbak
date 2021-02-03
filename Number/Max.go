package Number

import "math"

// Max 取出任意类型中数值类型的最大值,无数值类型则为0.
func (toolNumber *Number)Max(nums ...interface{}) (res float64) {
	if len(nums) < 1 {
		panic("[Max]: the nums length is less than 1")
	}

	var err error
	var val float64
	res, _ = toolNumber.ToFloat(nums[0])
	for _, v := range nums {
		val, err = toolNumber.ToFloat(v)
		if err == nil {
			res = math.Max(res, val)
		}
	}

	return
}
