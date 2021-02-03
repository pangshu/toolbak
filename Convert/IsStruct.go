package Convert

import "reflect"

// IsStruct 变量是否结构体.
func (*Convert)IsStruct(val interface{}) bool {
	r := reflect.ValueOf(val)
	// 如果是指针,则获取其所指向的元素
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	return r.Kind() == reflect.Struct
}
