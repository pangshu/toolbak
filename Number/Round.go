package Number

import "math"

// Round 对浮点数(的整数)进行四舍五入.
func (*Number)Round(value float64) float64 {
	return math.Floor(value + 0.5)
}
