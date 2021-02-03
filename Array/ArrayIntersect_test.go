package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestArrayIntersect Array/Global.go Array/ArrayIntersect.go Array/ArrayIntersect_test.go
func TestArrayIntersect(t *testing.T) {
	var toolArray Array
	arr1 := []string{"aa", "bb", "cc", "dd", "ee", "", "hh"}
	arr2 := []string{"bb", "cc", "ff", "gg", "ee", ""}
	res1 := toolArray.ArrayIntersect(arr1, arr2, "VALUE")
	fmt.Println(res1)
}

// go test -v -run TestArrayIntersect -bench=BenchmarkArrayIntersect -count=5 Array/Global.go Array/ArrayIntersect.go Array/ArrayIntersect_test.go
func BenchmarkArrayIntersect(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	arr1 := []string{"aa", "bb", "cc", "dd", "ee", "", "hh"}
	arr2 := []string{"bb", "cc", "ff", "gg", "ee", ""}
	for i:=0; i< t.N; i++ {
		_ = toolArray.ArrayIntersect(arr1, arr2)
	}
}