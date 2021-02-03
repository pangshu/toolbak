package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestInStringSlice Array/Global.go Array/InStringSlice.go Array/InStringSlice_test.go
func TestInStringSlice(t *testing.T) {
	var toolArray Array
	var arr1 = []string{"a","b","c","d","e","f"}
	res1 := toolArray.InStringSlice("d", arr1)
	fmt.Println(res1)
}

// go test -v -run TestInStringSlice -bench=BenchmarkInStringSlice -count=5 Array/Global.go Array/InStringSlice.go Array/InStringSlice_test.go
func BenchmarkInStringSlice(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	var arr1 = []string{"a","b","c","d","e","f"}
	for i:=0; i< t.N; i++ {
		_ = toolArray.InStringSlice("d",arr1)
	}
}