package Number

// Percent 返回百分比(val/total).
func (toolNumber *Number)Percent(val, total interface{}) float64 {
	t,_ := toolNumber.ToFloat(total)
	if t == 0 {
		return float64(0)
	}

	v,_ := toolNumber.ToFloat(val)

	return (v / t) * 100
}
