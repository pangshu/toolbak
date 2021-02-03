package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestUniqueInts Array/Global.go Array/UniqueInts.go Array/UniqueInts_test.go
func TestUniqueInts(t *testing.T) {
	var toolArray Array
	arr1 := []int{-3, 9, -5, 0, 5, -3, 0, 7}
	res1 := toolArray.UniqueInts(arr1)
	fmt.Println(res1)
}

// go test -v -run TestUniqueInts -bench=BenchmarkUniqueInts -count=5 Array/Global.go Array/UniqueInts.go Array/UniqueInts_test.go
func BenchmarkUniqueInts(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	arr1 := []int{-3, 9, -5, 0, 5, -3, 0, 7}
	for i:=0; i< t.N; i++ {
		_ = toolArray.UniqueInts(arr1)
	}
}