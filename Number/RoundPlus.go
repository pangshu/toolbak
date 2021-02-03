package Number

import "math"

// RoundPlus 对指定的小数位进行四舍五入.
// precision为小数位数.
func (toolNumber *Number)RoundPlus(value float64, precision int8) float64 {
	shift := math.Pow(10, float64(precision))
	return toolNumber.Round(value*shift) / shift
}
