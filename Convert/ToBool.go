package Convert

import "strconv"

// ToBool 强制将变量转换为布尔值.
func (toolConvert *Convert)ToBool(val interface{}) bool {
	switch val.(type) {
	case int:
		return (val.(int) > 0)
	case int8:
		return (val.(int8) > 0)
	case int16:
		return (val.(int16) > 0)
	case int32:
		return (val.(int32) > 0)
	case int64:
		return (val.(int64) > 0)
	case uint:
		return (val.(uint) > 0)
	case uint8:
		return (val.(uint8) > 0)
	case uint16:
		return (val.(uint16) > 0)
	case uint32:
		return (val.(uint32) > 0)
	case uint64:
		return (val.(uint64) > 0)
	case float32:
		return (val.(float32) > 0)
	case float64:
		return (val.(float64) > 0)
	case []uint8:
		res,_ := strconv.ParseBool(string(val.([]uint8)))
		return res
	case string:
		res,_ := strconv.ParseBool(val.(string))
		return res
	case bool:
		return val.(bool)
	default:
		return false
	}
}
