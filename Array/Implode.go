package Array

import (
	"bytes"
	"fmt"
	"reflect"
)

// Implode 用delimiter将数组(数组/切片/字典)的值连接为一个字符串.
func (*Array)Implode(delimiter string, arr interface{}) string {
	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		length := val.Len()
		if length == 0 {
			return ""
		}
		var buf bytes.Buffer
		j := length
		for i := 0; i < length; i++ {
			buf.WriteString(fmt.Sprintf("%s", val.Index(i)))
			if j--; j > 0 {
				buf.WriteString(delimiter)
			}
		}
		return buf.String()
	case reflect.Map:
		length := len(val.MapKeys())
		if length == 0 {
			return ""
		}
		var buf bytes.Buffer
		for _, k := range val.MapKeys() {
			buf.WriteString(fmt.Sprintf("%s", val.MapIndex(k).Interface()))
			if length--; length > 0 {
				buf.WriteString(delimiter)
			}
		}

		return buf.String()
	default:
		panic("[Implode]arr type must be array, slice")
	}
}
