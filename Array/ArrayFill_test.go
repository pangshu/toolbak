package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestArrayFill Array/Global.go Array/ArrayFill.go Array/ArrayFill_test.go
func TestArrayFill(t *testing.T) {
	var toolArray Array
	arr1 := "abcdefghi"
	num := 5
	res1 := toolArray.ArrayFill(arr1, num)
	fmt.Println(res1)
}

// go test -v -run TestArrayFill -bench=BenchmarkArrayFill -count=5 Array/Global.go Array/ArrayFill.go Array/ArrayFill_test.go
func BenchmarkArrayFill(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	arr1 := "abcdefghi"
	num := 5
	for i:=0; i< t.N; i++ {
		_ = toolArray.ArrayFill(arr1, num)
	}
}