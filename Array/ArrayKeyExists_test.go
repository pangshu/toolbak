package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestArrayKeyExists Array/Global.go Array/ArrayKeyExists.go Array/ArrayKeyExists_test.go
func TestArrayKeyExists(t *testing.T) {
	var toolArray Array
	arr1 := []string{"aa", "bb", "cc", "dd", "ee", "", "hh"}
	res1 := toolArray.ArrayKeyExists(arr1, 7)
	fmt.Println(res1)
}

// go test -v -run TestArrayKeyExists -bench=BenchmarkArrayKeyExists -count=5 Array/Global.go Array/ArrayKeyExists.go Array/ArrayKeyExists_test.go
func BenchmarkArrayKeyExists(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	arr1 := []string{"aa", "bb", "cc", "dd", "ee", "", "hh"}
	for i:=0; i< t.N; i++ {
		_ = toolArray.ArrayKeyExists(arr1, 1)
	}
}