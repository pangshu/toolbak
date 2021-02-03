package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestUniqueStrings Array/Global.go Array/UniqueStrings.go Array/UniqueStrings_test.go
func TestUniqueStrings(t *testing.T) {
	var toolArray Array
	arr1 := []string{"", "hello", "world", "hello", "你好", "world", "1234"}
	res1 := toolArray.UniqueStrings(arr1)
	fmt.Println(res1)
}

// go test -v -run TestUniqueStrings -bench=BenchmarkUniqueStrings -count=5 Array/Global.go Array/UniqueStrings.go Array/UniqueStrings_test.go
func BenchmarkUniqueStrings(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	arr1 := []string{"", "hello", "world", "hello", "你好", "world", "1234"}
	for i:=0; i< t.N; i++ {
		_ = toolArray.UniqueStrings(arr1)
	}
}