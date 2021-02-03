package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsMap Array/Global.go Array/IsMap.go Array/IsMap_test.go
func TestIsMap(t *testing.T) {
	var toolArray Array
	arr1 := []string{"a","b","c","d","e","f"}
	res1 := toolArray.IsMap(arr1)
	fmt.Println(res1)
}

// go test -v -run TestIsMap -bench=BenchmarkIsMap -count=5 Array/Global.go Array/IsMap.go Array/IsMap_test.go
func BenchmarkIsMap(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	arr1 := []string{"a","b","c","d","e","f"}
	for i:=0; i< t.N; i++ {
		_ = toolArray.IsMap(arr1)
	}
}