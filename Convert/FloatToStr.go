package Convert

import "strconv"

// FloatToStr 将浮点数转换为字符串,decimal为小数位数.
func (*Convert)FloatToStr(val interface{}, decimal int) string {
	switch val.(type) {
	// Floats
	case float32:
		return strconv.FormatFloat(float64(val.(float32)), 'f', decimal, 32)
	case float64:
		return strconv.FormatFloat(val.(float64), 'f', decimal, 64)
	// Type is not floats, return empty string
	default:
		return ""
	}
}
