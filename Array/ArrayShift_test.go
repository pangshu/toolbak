package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestArrayShift Array/Global.go Array/ArrayShift.go Array/ArrayShift_test.go
func TestArrayShift(t *testing.T) {
	var toolArray Array
	var arr1 = []interface{}{1,2,3,4,5,6,7,8,9,10}
	res1 := toolArray.ArrayShift(&arr1)
	fmt.Println(res1)
}

// go test -v -run TestArrayShift -bench=BenchmarkArrayShift -count=5 Array/Global.go Array/ArrayShift.go Array/ArrayShift_test.go
func BenchmarkArrayShift(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	var arr1 = []interface{}{1,2,3,4,5,6,7,8,9,10}
	for i:=0; i< t.N; i++ {
		_ = toolArray.ArrayShift(&arr1)
	}
}