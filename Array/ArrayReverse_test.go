package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestArrayReverse Array/Global.go Array/ArrayReverse.go Array/ArrayReverse_test.go
func TestArrayReverse(t *testing.T) {
	var toolArray Array
	var arr1 = []string{"aa", "bb", "cc", "dd", "ee", "ff", "hh"}
	res1 := toolArray.ArrayReverse(arr1)
	fmt.Println(res1)
}

// go test -v -run TestArrayReverse -bench=BenchmarkArrayReverse -count=5 Array/Global.go Array/ArrayReverse.go Array/ArrayReverse_test.go
func BenchmarkArrayReverse(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	var arr1 = []string{"aa", "bb", "cc", "dd", "ee", "ff", "hh"}
	for i:=0; i< t.N; i++ {
		_ = toolArray.ArrayReverse(arr1)
	}
}