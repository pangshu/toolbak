package Array

import (
	"fmt"
	"reflect"
)

// ArrayColumn 返回数组(切片/字典)中元素指定的一列.
// arr的元素必须是字典;
// columnKey为元素的字段名;
// 该方法效率较低.
func (*Array)ArrayColumn(arr interface{}, columnKey string) []interface{} {
	val := reflect.ValueOf(arr)
	var res []interface{}
	var item interface{}
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			item = val.Index(i).Interface()
			itemVal := reflect.ValueOf(item)
			switch itemVal.Kind() {
			case reflect.Map:
				for _, subKey := range itemVal.MapKeys() {
					if fmt.Sprintf("%s", subKey) == columnKey {
						res = append(res, itemVal.MapIndex(subKey).Interface())
						break
					}
				}
			default:
				panic("[ArrayColumn]arr`s slice item must be map")
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			item = val.MapIndex(k).Interface()
			itemVal := reflect.ValueOf(item)
			switch itemVal.Kind() {
			case reflect.Map:
				for _, subKey := range itemVal.MapKeys() {
					if fmt.Sprintf("%s", subKey) == columnKey {
						res = append(res, itemVal.MapIndex(subKey).Interface())
						break
					}
				}
			default:
				panic("[ArrayColumn]arr`s map item must be map")
			}
		}
	default:
		panic("[ArrayColumn]arr type must be array, slice or map")
	}

	return res
}
