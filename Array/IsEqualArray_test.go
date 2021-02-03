package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsEqualArray Array/Global.go Array/IsEqualArray.go Array/IsEqualArray_test.go
func TestIsEqualArray(t *testing.T) {
	var toolArray Array
	s1 := []string{"a", "b"}
	s2 := []string{"b", "a"}
	res1 := toolArray.IsEqualArray(s1, s2)
	fmt.Println(res1)
}

// go test -v -run TestIsEqualArray -bench=BenchmarkIsEqualArray -count=5 Array/Global.go Array/IsEqualArray.go Array/IsEqualArray_test.go
func BenchmarkIsEqualArray(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	s1 := []string{"a", "b"}
	s2 := []string{"b", "a"}
	for i:=0; i< t.N; i++ {
		_ = toolArray.IsEqualArray(s1, s2)
	}
}