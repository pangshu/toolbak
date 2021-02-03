package Array

import "reflect"

// CompareConditionMap 比对数组是否匹配条件.condition为条件字典,arr为要比对的数据数组.
func (*Array)CompareConditionMap(condition map[string]interface{}, arr interface{}) (res interface{}) {
	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Map:
		condLen := len(condition)
		chkNum := 0
		if condLen > 0 {
			for _, k := range val.MapKeys() {
				if condVal, ok := condition[k.String()]; ok && reflect.DeepEqual(val.MapIndex(k).Interface(), condVal) {
					chkNum++
				}
			}
		}

		if chkNum == condLen {
			res = arr
		}
	default:
		return
	}

	return
}
