package Array

import (
	"fmt"
	"reflect"
)

// MergeSlice 合并一个或多个数组/切片.
// filterNil是否过滤空元素(nil,''),true时排除空元素,false时保留空元素;ss是元素为数组/切片的数组.
func (toolArray *Array)MergeSlice(filterNil bool, ss ...interface{}) []interface{} {
	var res []interface{}
	switch len(ss) {
	case 0:
		break
	default:
		n := 0
		for i, v := range ss {
			chkLen := toolArray.IsArrayOrSlice(v, 3)
			if chkLen == -1 {
				msg := fmt.Sprintf("[MergeSlice]ss type must be array or slice, but [%d]th item not is.", i)
				panic(msg)
			} else {
				n += chkLen
			}
		}
		res = make([]interface{}, 0, n)
		var item interface{}
		for _, v := range ss {
			val := reflect.ValueOf(v)
			switch val.Kind() {
			case reflect.Array, reflect.Slice:
				for i := 0; i < val.Len(); i++ {
					item = val.Index(i).Interface()
					if !filterNil || (filterNil && item != nil && fmt.Sprintf("%v", item) != "") {
						res = append(res, item)
					}
				}
			}
		}
	}
	return res
}
