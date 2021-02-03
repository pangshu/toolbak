package Array

import (
	"math/rand"
	"reflect"
	"time"
)

// ArrayRand 从数组/切片中随机取出num个单元.
func (*Array)ArrayRand(arr interface{}, num int) []interface{} {
	if num < 1 {
		panic("[ArrayRand]num: cannot be less than 1")
	}

	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		length := val.Len()
		if length == 0 {
			return nil
		}
		if num > length {
			num = length
		}
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		res := make([]interface{}, num)
		for i, v := range r.Perm(length) {
			if i < num {
				res[i] = val.Index(v).Interface()
			} else {
				break
			}
		}
		return res
	default:
		//panic("[ArrayRand]arr type must be array or slice")
		return nil
	}
}
