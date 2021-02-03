package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestJoinStrings Array/Global.go Array/JoinStrings.go Array/JoinStrings_test.go
func TestJoinStrings(t *testing.T) {
	var toolArray Array
	arr1 := []string{"a","b","c","d","e","f"}
	res1 := toolArray.JoinStrings(arr1, ",")
	fmt.Println(res1)
}

// go test -v -run TestJoinStrings -bench=BenchmarkJoinStrings -count=5 Array/Global.go Array/JoinStrings.go Array/JoinStrings_test.go
func BenchmarkJoinStrings(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	arr1 := []string{"a","b","c","d","e","f"}
	for i:=0; i< t.N; i++ {
		_ = toolArray.JoinStrings(arr1, ",")
	}
}