package Number

import (
	"reflect"
	"toolbak/Convert"
)

// InRange 数值是否在某个范围内,将自动转换类型再比较.
func (toolNumber *Number)InRange(value interface{}, left interface{}, right interface{}) bool {
	var toolConvert Convert.Convert
	reflectValue := reflect.TypeOf(value).Kind()
	reflectLeft := reflect.TypeOf(left).Kind()
	reflectRight := reflect.TypeOf(right).Kind()

	if reflectValue == reflect.Int && reflectLeft == reflect.Int && reflectRight == reflect.Int {
		return toolNumber.InRangeInt(value.(int), left.(int), right.(int))
	} else if reflectValue == reflect.Float32 && reflectLeft == reflect.Float32 && reflectRight == reflect.Float32 {
		return toolNumber.InRangeFloat32(value.(float32), left.(float32), right.(float32))
	} else if reflectValue == reflect.Float64 && reflectLeft == reflect.Float64 && reflectRight == reflect.Float64 {
		return toolNumber.InRangeFloat64(value.(float64), left.(float64), right.(float64))
	} else if toolConvert.IsInt(value) && toolConvert.IsInt(left) && toolConvert.IsInt(right) {
		return toolNumber.InRangeInt(toolConvert.ToInt(value), toolConvert.ToInt(left), toolConvert.ToInt(right))
	} else if toolConvert.IsNumeric(value) && toolConvert.IsNumeric(left) && toolConvert.IsNumeric(right) {
		return toolNumber.InRangeFloat64(toolConvert.ToFloat(value), toolConvert.ToFloat(left), toolConvert.ToFloat(right))
	}

	return false
}
