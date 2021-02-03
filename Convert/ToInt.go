package Convert

import (
	"strconv"
	"strings"
)

// ToInt 强制将变量转换为整型;其中true或"true"为1.
func (toolConvert *Convert)ToInt(val interface{}) int {
	switch val.(type) {
	case int:
		return val.(int)
	case int8:
		return int(val.(int8))
	case int16:
		return int(val.(int16))
	case int32:
		return int(val.(int32))
	case int64:
		return int(val.(int64))
	case uint:
		return int(val.(uint))
	case uint8:
		return int(val.(uint8))
	case uint16:
		return int(val.(uint16))
	case uint32:
		return int(val.(uint32))
	case uint64:
		return int(val.(uint64))
	case float32:
		return int(val.(float32))
	case float64:
		return int(val.(float64))
	case []uint8:
		if strings.ToLower(val.(string)) == "true" {
			return 1
		} else {
			res, _ := strconv.Atoi(val.(string))
			return res
		}
	case string:
		if strings.ToLower(val.(string)) == "true" {
			return 1
		} else {
			res, _ := strconv.Atoi(val.(string))
			return res
		}
	case bool:
		if val.(bool) == true {
			return 1
		} else {
			return 0
		}
	default:
		return 0
	}
}
