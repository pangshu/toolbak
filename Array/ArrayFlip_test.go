package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestArrayFlip Array/Global.go Array/ArrayFlip.go Array/ArrayFlip_test.go
func TestArrayFlip(t *testing.T) {
	var toolArray Array
	arr1 := map[string]int{"a": 1, "b": 2, "c": 3}
	res1 := toolArray.ArrayFlip(arr1)
	fmt.Println(res1)
}

// go test -v -run TestArrayFlip -bench=BenchmarkArrayFlip -count=5 Array/Global.go Array/ArrayFlip.go Array/ArrayFlip_test.go
func BenchmarkArrayFlip(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	arr1 := map[string]int{"a": 1, "b": 2, "c": 3}
	for i:=0; i< t.N; i++ {
		_ = toolArray.ArrayFlip(arr1)
	}
}