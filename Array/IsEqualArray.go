package Array

import (
	"fmt"
	"reflect"
)

// IsEqualArray 两个数组/切片是否相同(不管元素顺序).
// expected, actual 是要比较的数组/切片.
func (*Array)IsEqualArray(expected, actual interface{}) bool {
	var itmAsStr string

	expectedMap := map[string]bool{}
	val := reflect.ValueOf(expected)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			itmAsStr = fmt.Sprintf("%#v", val.Index(i).Interface())
			expectedMap[itmAsStr] = true
		}
	default:
		panic("[IsEqualArray]expected type must be array or slice")
	}

	actualMap := map[string]bool{}
	val = reflect.ValueOf(actual)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			itmAsStr = fmt.Sprintf("%#v", val.Index(i).Interface())
			actualMap[itmAsStr] = true
		}
	default:
		panic("[IsEqualArray]actual type must be array or slice")
	}

	return reflect.DeepEqual(expectedMap, actualMap)
}
