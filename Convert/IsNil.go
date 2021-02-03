package Convert

import "reflect"

// IsNil 检查变量是否空值.
func (*Convert)IsNil(val interface{}) bool {
	if val == nil {
		return true
	}

	rv := reflect.ValueOf(val)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice, reflect.Interface:
		if rv.IsNil() {
			return true
		}
	}
	return false
}
