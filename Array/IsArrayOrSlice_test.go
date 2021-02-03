package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsArrayOrSlice Array/Global.go Array/IsArrayOrSlice.go Array/IsArrayOrSlice_test.go
func TestIsArrayOrSlice(t *testing.T) {
	var toolArray Array
	var arr1 = []int{1, 2, 3, 4, 5, 6}
	res1 := toolArray.IsArrayOrSlice(arr1, 2)
	fmt.Println(res1)
}

// go test -v -run TestIsArrayOrSlice -bench=BenchmarkIsArrayOrSlice -count=5 Array/Global.go Array/IsArrayOrSlice.go Array/IsArrayOrSlice_test.go
func BenchmarkIsArrayOrSlice(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	var arr1 = []string{"a","b","c","d","e","f"}
	for i:=0; i< t.N; i++ {
		_ = toolArray.IsArrayOrSlice(arr1, 1)
	}
}