package Array

import (
	"fmt"
	"reflect"
)

// ArrayKeyExists 检查数组里是否有指定的键名或索引.
func (*Array)ArrayKeyExists(arr interface{}, key interface{}) bool {
	if key == nil {
		return false
	}

	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		var keyInt int
		var keyIsInt, ok bool
		if keyInt, ok = key.(int); ok {
			keyIsInt = true
		}

		length := val.Len()
		if keyIsInt && length > 0 && keyInt >= 0 && keyInt < length {
			return true
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if fmt.Sprintf("%s", key) == fmt.Sprintf("%s", k) || reflect.DeepEqual(key, k) {
				return true
			}
		}
	default:
		panic("[ArrayKeyExists]arr type must be array, slice or map")
	}
	return false
}
