package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestImplode Array/Global.go Array/Implode.go Array/Implode_test.go
func TestImplode(t *testing.T) {
	var toolArray Array
	var arr1 = []string{"a", "b", "c", "d", "e"}
	res1 := toolArray.Implode(",", arr1)
	fmt.Println(res1)
}

// go test -v -run TestImplode -bench=BenchmarkImplode -count=5 Array/Global.go Array/Implode.go Array/Implode_test.go
func BenchmarkImplode(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	var arr1 = []string{"a", "b", "c", "d", "e"}
	for i:=0; i< t.N; i++ {
		_ = toolArray.Implode(",",arr1)
	}
}