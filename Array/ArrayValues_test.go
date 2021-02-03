package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestArrayValues Array/Global.go Array/ArrayValues.go Array/ArrayValues_test.go
func TestArrayValues(t *testing.T) {
	var toolArray Array
	var arr1 = []interface{}{1,2,3,4,5,6,7,8,9,1}
	res1 := toolArray.ArrayValues(arr1,false)
	fmt.Println(res1)
}

// go test -v -run TestArrayValues -bench=BenchmarkArrayValues -count=5 Array/Global.go Array/ArrayValues.go Array/ArrayValues_test.go
func BenchmarkArrayValues(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	var arr1 = []interface{}{1,2,3,4,5,6,7,8,9,1}
	for i:=0; i< t.N; i++ {
		_ = toolArray.ArrayValues(arr1,false)
	}
}