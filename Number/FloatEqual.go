package Number

import "math"

// FloatEqual 比较两个浮点数是否相等.decimal为小数精确位数.
func (*Number)FloatEqual(f1 float64, f2 float64, decimal ...int) bool {
	var threshold float64
	var dec int
	if len(decimal) == 0 {
		dec = 10
	} else {
		dec = decimal[0]
	}

	//比较精度
	threshold = math.Pow10(-dec)

	return math.Abs(f1-f2) <= threshold
}

