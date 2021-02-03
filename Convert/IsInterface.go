package Convert

import "reflect"

// IsInterface 变量是否接口.
func (*Convert)IsInterface(val interface{}) bool {
	// 如果是指针,则获取其所指向的元素
	r := reflect.ValueOf(val)
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	return r.Kind() == reflect.Invalid
}
