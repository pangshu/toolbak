package Array

import "reflect"

// ArraySearchMutil 从数组中搜索对应元素(多个).
// arr为要查找的数组,元素必须为字典;condition为条件字典.
func (toolArray *Array)ArraySearchMutil(arr interface{}, condition map[string]interface{}) (res []interface{}) {
	// 条件为空
	if len(condition) == 0 {
		return
	}

	val := reflect.ValueOf(arr)
	var item interface{}
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			item = toolArray.CompareConditionMap(condition, val.Index(i).Interface())
			if item != nil {
				res = append(res, item)
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			item = toolArray.CompareConditionMap(condition, val.MapIndex(k).Interface())
			if item != nil {
				res = append(res, item)
			}
		}
	default:
		panic("[ArraySearchMutil]arr type must be array, slice or map")
	}

	return
}
