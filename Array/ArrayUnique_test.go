package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestArrayUnique Array/Global.go Array/ArrayUnique.go Array/ArrayUnique_test.go
func TestArrayUnique(t *testing.T) {
	var toolArray Array
	var arr1 = []interface{}{1,2,3,4,5,6,7,8,9,1}
	res1 := toolArray.ArrayUnique(arr1)
	fmt.Println(res1)
}

// go test -v -run TestArrayUnique -bench=BenchmarkArrayUnique -count=5 Array/Global.go Array/ArrayUnique.go Array/ArrayUnique_test.go
func BenchmarkArrayUnique(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	var arr1 = []interface{}{1,2,3,4,5,6,7,8,9,1}
	for i:=0; i< t.N; i++ {
		_ = toolArray.ArrayUnique(arr1)
	}
}