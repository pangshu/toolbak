package Array

import "reflect"

// ArraySlice 返回根据 offset 和 size 参数所指定的 arr 数组/切片中的一段切片.
func (*Array)ArraySlice(arr interface{}, offset, size int) []interface{} {
	if size < 1 {
		return nil
		//panic("[ArraySlice]size: cannot be less than 1")
	}

	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		length := val.Len()
		if length == 0 || (offset > 0 && offset > length-1) {
			return nil
		}

		items := make([]interface{}, length)
		for i := 0; i < val.Len(); i++ {
			items[i] = val.Index(i).Interface()
		}

		if offset < 0 {
			offset = offset%length + length
		}
		end := offset + size
		if end < length {
			return items[offset:end]
		}
		return items[offset:]
	default:
		//panic("[ArraySlice]arr type must be array or slice")
		return nil
	}
}
