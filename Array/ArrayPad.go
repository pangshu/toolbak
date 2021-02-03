package Array

import "reflect"

// ArrayPad 以指定长度将一个值item填充进arr数组/切片.
// 若 size 为正，则填补到数组的右侧，如果为负则从左侧开始填补;
// 若 size 的绝对值小于或等于 arr 数组的长度则没有任何填补.
func (toolArray *Array)ArrayPad(arr interface{}, size int, item interface{}) []interface{} {
	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		length := val.Len()
		if length == 0 && size > 0 {
			return toolArray.ArrayFill(item, size)
		}

		orig := make([]interface{}, length)
		for i := 0; i < length; i++ {
			orig[i] = val.Index(i).Interface()
		}

		if size == 0 || (size > 0 && size < length) || (size < 0 && size > -length) {
			return orig
		}

		n := size
		if size < 0 {
			n = -size
		}
		n -= length
		items := make([]interface{}, n)
		for i := 0; i < n; i++ {
			items[i] = item
		}

		if size > 0 {
			return append(orig, items...)
		}
		return append(items, orig...)
	default:
		panic("[ArrayPad]arr type must be array, slice")
	}
}
