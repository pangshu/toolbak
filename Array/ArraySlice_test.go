package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestArraySlice Array/Global.go Array/ArraySlice.go Array/ArraySlice_test.go
func TestArraySlice(t *testing.T) {
	var toolArray Array
	var arr1 = []interface{}{1,2,3,4,5,6,7,8,9,10}
	res1 := toolArray.ArraySlice(arr1,0,1)
	fmt.Println(res1)
}

// go test -v -run TestArraySlice -bench=BenchmarkArraySlice -count=5 Array/Global.go Array/ArraySlice.go Array/ArraySlice_test.go
func BenchmarkArraySlice(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	var arr1 = []interface{}{1,2,3,4,5,6,7,8,9,10}
	for i:=0; i< t.N; i++ {
		_ = toolArray.ArraySlice(arr1,0,1)
	}
}