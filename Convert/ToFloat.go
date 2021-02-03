package Convert

// ToFloat 强制将变量转换为浮点型;其中true或"true"为1.0 .
func (toolConvert *Convert)ToFloat(val interface{}) (res float64) {
	switch val.(type) {
	case int:
		res = float64(val.(int))
	case int8:
		res = float64(val.(int8))
	case int16:
		res = float64(val.(int16))
	case int32:
		res = float64(val.(int32))
	case int64:
		res = float64(val.(int64))
	case uint:
		res = float64(val.(uint))
	case uint8:
		res = float64(val.(uint8))
	case uint16:
		res = float64(val.(uint16))
	case uint32:
		res = float64(val.(uint32))
	case uint64:
		res = float64(val.(uint64))
	case float32:
		res = float64(val.(float32))
	case float64:
		res = val.(float64)
	case []uint8:
		res = toolConvert.StrToFloat64(string(val.([]uint8)))
	case string:
		res = toolConvert.StrToFloat64(val.(string))
	case bool:
		if val.(bool) {
			res = 1.0
		}
	}

	return
}
