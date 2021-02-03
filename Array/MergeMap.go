package Array

import (
	"fmt"
	"reflect"
)

// MergeMap 合并字典.
// 相同的键名时,后面的值将覆盖前一个值;key2Str是否将键转换为字符串;ss是元素为字典的数组.
func (*Array)MergeMap(key2Str bool, ss ...interface{}) map[interface{}]interface{} {
	res := make(map[interface{}]interface{})
	switch len(ss) {
	case 0:
		break
	default:
		for i, v := range ss {
			val := reflect.ValueOf(v)
			switch val.Kind() {
			case reflect.Map:
				for _, k := range val.MapKeys() {
					if key2Str {
						res[k.String()] = val.MapIndex(k).Interface()
					} else {
						res[k] = val.MapIndex(k).Interface()
					}
				}
			default:
				msg := fmt.Sprintf("[MergeMap]ss type must be map, but [%d]th item not is.", i)
				panic(msg)
			}
		}
	}
	return res
}
