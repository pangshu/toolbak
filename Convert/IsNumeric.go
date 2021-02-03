package Convert

import "strconv"

// IsNumeric 变量是否数值(不包含复数和科学计数法).
func (*Convert)IsNumeric(val interface{}) bool {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	case float32, float64:
		return true
	case string:
		str := val.(string)
		if str == "" {
			return false
		}
		_, err := strconv.ParseFloat(str, 64)
		return err == nil
	}

	return false
}
