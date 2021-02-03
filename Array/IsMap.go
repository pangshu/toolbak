package Array

import "reflect"

// IsMap 检查变量是否字典.
func (*Array)IsMap(val interface{}) bool {
	return reflect.ValueOf(val).Kind() == reflect.Map
}
