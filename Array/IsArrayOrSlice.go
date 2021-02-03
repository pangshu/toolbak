package Array

import (
	"fmt"
	"reflect"
)

// IsArrayOrSlice 检查变量是否数组或切片.
// chkType检查类型,枚举值有(1仅数组,2仅切片,3数组或切片);结果为-1表示非,>=0表示是
func (*Array)IsArrayOrSlice(data interface{}, chkType uint8) int {
	if chkType != 1 && chkType != 2 && chkType != 3 {
		panic(fmt.Sprintf("[isArrayOrSlice] chkType value muset in (1, 2, 3), but it`s %d", chkType))
	}

	var res = -1
	val := reflect.ValueOf(data)
	switch val.Kind() {
	case reflect.Array:
		if chkType == 1 || chkType == 3 {
			res = val.Len()
		}
	case reflect.Slice:
		if chkType == 2 || chkType == 3 {
			res = val.Len()
		}
	}

	return res
}
