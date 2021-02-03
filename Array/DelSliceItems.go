package Array

import "reflect"

// DelSliceItems 删除数组/切片的元素,返回一个新切片.
// ids为多个元素的索引(0~len(val)-1);
// del为删除元素的数量.
func (*Array)DelSliceItems(val interface{}, ids ...int) (res []interface{}, del int) {
	sl := reflect.ValueOf(val)
	styp := sl.Kind()

	if styp != reflect.Array && styp != reflect.Slice {
		panic("[DeleteSlice] val type must be array or slice")
	}

	slen := sl.Len()
	if slen == 0 {
		return
	}

	var item interface{}
	var chk bool
	for i := 0; i < slen; i++ {
		item = sl.Index(i).Interface()
		chk = true

		for _, v := range ids {
			if i == v {
				del++
				chk = false
				break
			}
		}

		if chk {
			res = append(res, item)
		}
	}

	return
}
