package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestInIntSlice Array/Global.go Array/InIntSlice.go Array/InIntSlice_test.go
func TestInIntSlice(t *testing.T) {
	var toolArray Array
	var arr1 = []int{1,2,3,4,5,6,7}
	res1 := toolArray.InIntSlice(5, arr1)
	fmt.Println(res1)
}

// go test -v -run TestInIntSlice -bench=BenchmarkInIntSlice -count=5 Array/Global.go Array/InIntSlice.go Array/InIntSlice_test.go
func BenchmarkInIntSlice(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	var arr1 = []int{1,2,3,4,5,6,7}
	for i:=0; i< t.N; i++ {
		_ = toolArray.InIntSlice(5,arr1)
	}
}