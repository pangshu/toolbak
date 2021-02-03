package Array

import (
	"fmt"
	"reflect"
)

// ArraySearchItem 从数组中搜索对应元素(单个).
// arr为要查找的数组,元素必须为字典;condition为条件字典.
func (toolArray *Array)ArraySearchItem(arr interface{}, condition map[string]interface{}) (res interface{}) {
	// 条件为空
	if len(condition) == 0 {
		return
	}

	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		fmt.Println("-----array---slice------")
		for i := 0; i < val.Len(); i++ {
			res = toolArray.CompareConditionMap(condition, val.Index(i).Interface())
			if res != nil {
				return
			}
		}
	case reflect.Map:
		fmt.Println("-----map------")
		fmt.Println(val.MapKeys())
		for _, k := range val.MapKeys() {
			fmt.Println(val.MapIndex(k).Interface())
			res = toolArray.CompareConditionMap(condition, val.MapIndex(k).Interface())
			if res != nil {
				return
			}
		}
	default:
		//panic("[ArraySearchItem]arr type must be array, slice or map")
		return nil
	}

	return
}
