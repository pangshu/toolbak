package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestInInt64Slice Array/Global.go Array/InInt64Slice.go Array/InInt64Slice_test.go
func TestInInt64Slice(t *testing.T) {
	var toolArray Array
	var arr1 = []int64{1,2,3,4,5,6,7}
	res1 := toolArray.InInt64Slice(5, arr1)
	fmt.Println(res1)
}

// go test -v -run TestInInt64Slice -bench=BenchmarkInInt64Slice -count=5 Array/Global.go Array/InInt64Slice.go Array/InInt64Slice_test.go
func BenchmarkInInt64Slice(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	var arr1 = []int64{1,2,3,4,5,6,7}
	for i:=0; i< t.N; i++ {
		_ = toolArray.InInt64Slice(5,arr1)
	}
}