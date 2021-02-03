package Convert

import "regexp"

// IsFloat 变量是否浮点数值.
func (*Convert)IsFloat(val interface{}) bool {
	switch val.(type) {
	case float32, float64:
		return true
	case string:
		str := val.(string)
		if str == "" {
			return false
		}

		if ok := regexp.MustCompile(`^(-?\d+)(\.\d+)?`).MatchString(str); ok {
			return true
		}
	}

	return false
}
