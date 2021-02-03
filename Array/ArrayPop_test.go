package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestArrayPop Array/Global.go Array/ArrayPop.go Array/ArrayPop_test.go Array/ArrayFill.go
func TestArrayPop(t *testing.T) {
	var toolArray Array
	//arr1 := []string{"aa", "bb", "cc", "dd", "ee", "ff", "hh"}
	var arr1 = []interface{}{"aa", "bb", "cc", "dd", "ee", "ff", "hh"}

	res1 := toolArray.ArrayPop(&arr1)
	fmt.Println(res1)
}

// go test -v -run TestArrayPop -bench=BenchmarkArrayPop -count=5 Array/Global.go Array/ArrayPop.go Array/ArrayPop_test.go
func BenchmarkArrayPop(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	arr1 := []interface{}{"aa", "bb", "cc", "dd", "ee", "", "hh"}
	for i:=0; i< t.N; i++ {
		_ = toolArray.ArrayPop(&arr1)
	}
}