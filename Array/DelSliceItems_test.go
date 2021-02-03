package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestDelSliceItems Array/Global.go Array/DelSliceItems.go Array/DelSliceItems_test.go
func TestDelSliceItems(t *testing.T) {
	var toolArray Array
	var arr1 = []interface{}{1,2,3,4,5,6,7,8,9,1}
	res1,res2 := toolArray.DelSliceItems(arr1,0,1,2)
	fmt.Println(res1)
	fmt.Println(res2)
}

// go test -v -run TestDelSliceItems -bench=BenchmarkDelSliceItems -count=5 Array/Global.go Array/DelSliceItems.go Array/DelSliceItems_test.go
func BenchmarkDelSliceItems(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	var arr1 = []interface{}{1,2,3,4,5,6,7,8,9,1}
	for i:=0; i< t.N; i++ {
		_,_ = toolArray.DelSliceItems(arr1,2)
	}
}