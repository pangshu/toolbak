package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestArrayPush Array/Global.go Array/ArrayPush.go Array/ArrayPush_test.go Array/ArrayFill.go
func TestArrayPush(t *testing.T) {
	var toolArray Array
	var arr1 []interface{}
	res1 := toolArray.ArrayPush(&arr1, 1, 2, 3, "a", "b", "c")
	fmt.Println(res1)
}

// go test -v -run TestArrayPush -bench=BenchmarkArrayPush -count=5 Array/Global.go Array/ArrayPush.go Array/ArrayPush_test.go
func BenchmarkArrayPush(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	var arr1 []interface{}
	for i:=0; i< t.N; i++ {
		_ = toolArray.ArrayPush(&arr1, 1, 2, 3, "a", "b", "c")
	}
}