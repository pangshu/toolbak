package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestMergeMap Array/Global.go Array/MergeMap.go Array/MergeMap_test.go
func TestMergeMap(t *testing.T) {
	var toolArray Array
	arr1 := map[string]string{"a":"a","b":"b","c":"c","d":"d","e":"e","f":"f"}
	arr2 := map[string]string{"d":"d","e":"e","f":"f","g":"g","h":"h","i":"i"}
	res1 := toolArray.MergeMap(true,arr1, arr2)
	fmt.Println(res1)
}

// go test -v -run TestMergeMap -bench=BenchmarkMergeMap -count=5 Array/Global.go Array/MergeMap.go Array/MergeMap_test.go
func BenchmarkMergeMap(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	arr1 := map[string]string{"a":"a","b":"b","c":"c","d":"d","e":"e","f":"f"}
	arr2 := map[string]string{"d":"d","e":"e","f":"f","g":"g","h":"h","i":"i"}
	for i:=0; i< t.N; i++ {
		_ = toolArray.MergeMap(true,arr1,arr2)
	}
}