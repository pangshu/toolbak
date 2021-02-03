package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestArrayKeys Array/Global.go Array/ArrayKeys.go Array/ArrayKeys_test.go
func TestArrayKeys(t *testing.T) {
	var toolArray Array
	arr1 := []string{"aa", "bb", "cc", "dd", "ee", "", "hh"}
	res1 := toolArray.ArrayKeys(arr1)
	fmt.Println(res1)
}

// go test -v -run TestArrayKeys -bench=BenchmarkArrayKeys -count=5 Array/Global.go Array/ArrayKeys.go Array/ArrayKeys_test.go
func BenchmarkArrayKeys(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	arr1 := []string{"aa", "bb", "cc", "dd", "ee", "", "hh"}
	for i:=0; i< t.N; i++ {
		_ = toolArray.ArrayKeys(arr1)
	}
}