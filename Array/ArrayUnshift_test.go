package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestArrayUnshift Array/Global.go Array/ArrayUnshift.go Array/ArrayUnshift_test.go
func TestArrayUnshift(t *testing.T) {
	var toolArray Array
	var arr1 = []interface{}{1,2,3,4,5,6,7,8,9,1}
	res1 := toolArray.ArrayUnshift(&arr1,11)
	fmt.Println(res1)
}

// go test -v -run TestArrayUnshift -bench=BenchmarkArrayUnshift -count=5 Array/Global.go Array/ArrayUnshift.go Array/ArrayUnshift_test.go
func BenchmarkArrayUnshift(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	var arr1 = []interface{}{1,2,3,4,5,6,7,8,9,1}
	for i:=0; i< t.N; i++ {
		_ = toolArray.ArrayUnshift(&arr1,11)
	}
}