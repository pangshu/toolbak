package Array

import (
	"fmt"
	"reflect"
)

// ArrayValues 返回数组(切片/字典)中所有的值.
// filterNil是否过滤空元素(nil,''),true时排除空元素,false时保留空元素.
func (*Array)ArrayValues(arr interface{}, filterNil bool) []interface{} {
	var res []interface{}
	var item interface{}
	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			item = val.Index(i).Interface()
			if !filterNil || (filterNil && item != nil && fmt.Sprintf("%v", item) != "") {
				res = append(res, item)
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			item = val.MapIndex(k).Interface()
			if !filterNil || (filterNil && item != nil && fmt.Sprintf("%v", item) != "") {
				res = append(res, item)
			}
		}
	default:
		panic("[arrayValues] arr type must be array, slice or map")
	}

	return res
}
