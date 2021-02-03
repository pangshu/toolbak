package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestUnique64Ints Array/Global.go Array/Unique64Ints.go Array/Unique64Ints_test.go
func TestUnique64Ints(t *testing.T) {
	var toolArray Array
	arr1 := []int64{-3, 9, -5, 0, 5, -3, 0, 7}
	res1 := toolArray.Unique64Ints(arr1)
	fmt.Println(res1)
}

// go test -v -run TestUnique64Ints -bench=BenchmarkUnique64Ints -count=5 Array/Global.go Array/Unique64Ints.go Array/Unique64Ints_test.go
func BenchmarkUnique64Ints(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	arr1 := []int64{-3, 9, -5, 0, 5, -3, 0, 7}
	for i:=0; i< t.N; i++ {
		_ = toolArray.Unique64Ints(arr1)
	}
}