package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestMergeSlice Array/Global.go Array/MergeSlice.go Array/MergeSlice_test.go Array/IsArrayOrSlice.go
func TestMergeSlice(t *testing.T) {
	var toolArray Array
	arr1 := []string{"a","b","c","d","e","f"}
	arr2 := []string{"d","e","f","g","h","i"}
	res1 := toolArray.MergeSlice(true,arr1, arr2)
	fmt.Println(res1)
}

// go test -v -run TestMergeSlice -bench=BenchmarkMergeSlice -count=5 Array/Global.go Array/MergeSlice.go Array/MergeSlice_test.go Array/IsArrayOrSlice.go
func BenchmarkMergeSlice(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	arr1 := []string{"a","b","c","d","e","f"}
	arr2 := []string{"d","e","f","g","h","i"}
	for i:=0; i< t.N; i++ {
		_ = toolArray.MergeSlice(true,arr1,arr2)
	}
}