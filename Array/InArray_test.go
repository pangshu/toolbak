package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestInArray Array/Global.go Array/InArray.go Array/InArray_test.go
func TestInArray(t *testing.T) {
	var toolArray Array
	var arr1 = []string{"a", "b", "c", "d", "e"}
	res1 := toolArray.InArray("e", arr1)
	fmt.Println(res1)
}

// go test -v -run TestInArray -bench=BenchmarkInArray -count=5 Array/Global.go Array/InArray.go Array/InArray_test.go
func BenchmarkInArray(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	var arr1 = []string{"a", "b", "c", "d", "e"}
	for i:=0; i< t.N; i++ {
		_ = toolArray.InArray("e",arr1)
	}
}