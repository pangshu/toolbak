package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestArrayRand Array/Global.go Array/ArrayRand.go Array/ArrayRand_test.go
func TestArrayRand(t *testing.T) {
	var toolArray Array
	var arr1 = []string{"aa", "bb", "cc", "dd", "ee", "ff", "hh"}
	res1 := toolArray.ArrayRand(arr1, 5)
	fmt.Println(res1)
}

// go test -v -run TestArrayRand -bench=BenchmarkArrayRand -count=5 Array/Global.go Array/ArrayRand.go Array/ArrayRand_test.go
func BenchmarkArrayRand(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	var arr1 = []string{"aa", "bb", "cc", "dd", "ee", "ff", "hh"}
	for i:=0; i< t.N; i++ {
		_ = toolArray.ArrayRand(arr1, 1)
	}
}