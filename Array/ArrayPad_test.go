package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestArrayPad Array/Global.go Array/ArrayPad.go Array/ArrayPad_test.go Array/ArrayFill.go
func TestArrayPad(t *testing.T) {
	var toolArray Array
	//arr1 := []string{"aa", "bb", "cc", "dd", "ee", "ff", "hh"}
	var arr1 = []string{"aa", "bb", "cc", "dd", "ee", "ff", "hh"}

	res1 := toolArray.ArrayPad(arr1, 10, 10)
	fmt.Println(res1)
}

// go test -v -run TestArrayPad -bench=BenchmarkArrayPad -count=5 Array/Global.go Array/ArrayPad.go Array/ArrayPad_test.go
func BenchmarkArrayPad(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	arr1 := []string{"aa", "bb", "cc", "dd", "ee", "", "hh"}
	for i:=0; i< t.N; i++ {
		_ = toolArray.ArrayPad(arr1,5,5)
	}
}