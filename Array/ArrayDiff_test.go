package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestArrayDiff Array/Global.go Array/ArrayDiff.go Array/ArrayDiff_test.go
func TestArrayDiff(t *testing.T) {
	var toolArray Array
	arr1 := []string{"aa", "bb", "cc", "dd", "ee", "", "hh"}
	arr2 := []string{"bb", "cc", "ff", "gg", "ee", ""}
	res1 := toolArray.ArrayDiff(arr1, arr2, "VALUE")
	fmt.Println(res1)
}

// go test -v -run TestArrayDiff -bench=BenchmarkArrayDiff -count=5 Array/Global.go Array/ArrayDiff.go Array/ArrayDiff_test.go
func BenchmarkArrayDiff(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	arr1 := []string{"aa", "bb", "cc", "dd", "ee", "", "hh"}
	arr2 := []string{"bb", "cc", "ff", "gg", "ee", ""}
	for i:=0; i< t.N; i++ {
		_ = toolArray.ArrayDiff(arr1, arr2)
	}
}