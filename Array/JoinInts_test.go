package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestJoinInts Array/Global.go Array/JoinInts.go Array/JoinInts_test.go
func TestJoinInts(t *testing.T) {
	var toolArray Array
	arr1 := []int{1,2,3,4,5,6,7}
	res1 := toolArray.JoinInts(arr1,">")
	fmt.Println(res1)
}

// go test -v -run TestJoinInts -bench=BenchmarkJoinInts -count=5 Array/Global.go Array/JoinInts.go Array/JoinInts_test.go
func BenchmarkJoinInts(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	arr1 := []int{1,2,3,4,5,6,7}
	for i:=0; i< t.N; i++ {
		_ = toolArray.JoinInts(arr1, ">")
	}
}