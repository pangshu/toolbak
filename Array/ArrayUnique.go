package Array

import (
	"crypto/md5"
	"fmt"
	"reflect"
)

// ArrayUnique 移除数组中重复的值.
func (*Array)ArrayUnique(arr interface{}) []interface{} {
	val := reflect.ValueOf(arr)
	var res []interface{}
	var item interface{}
	var str, key string
	mp := make(map[string]interface{})
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			item = val.Index(i).Interface()
			str = fmt.Sprintf("%+v", item)
			key = fmt.Sprintf("%x", md5.Sum([]byte(str)))
			if _, ok := mp[key]; !ok {
				mp[key] = true
				res = append(res, item)
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			item = val.MapIndex(k).Interface()
			str = fmt.Sprintf("%+v", item)
			key = fmt.Sprintf("%x", md5.Sum([]byte(str)))
			if _, ok := mp[key]; !ok {
				mp[key] = true
				res = append(res, item)
			}
		}
	default:
		//panic("[ArrayUnique]arr type must be array, slice or map")
		return nil
	}

	return res
}
