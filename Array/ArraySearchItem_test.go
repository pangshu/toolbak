package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestArraySearchItem Array/Global.go Array/ArraySearchItem.go Array/ArraySearchItem_test.go Array/CompareConditionMap.go
func TestArraySearchItem(t *testing.T) {
	var toolArray Array
	var list []interface{}
	type item map[string]interface{}
	item1 := item{"age": 20, "name": "test1"}
	item2 := item{"age": 21, "name": "test2"}
	list = append(list, item1, item2, nil, "hello")

	//var arr1 = map[string]interface{}{"aa":"ab", "bb":"bc", "cc":"cd", "dd":"de", "ee":"ef", "ff":"fh", "hh":"hg"}
	var arr2 = map[string]interface{}{"age":21}
	res1 := toolArray.ArraySearchItem(list, arr2)
	fmt.Println(res1)
}

// go test -v -run TestArraySearchItem -bench=BenchmarkArraySearchItem -count=5 Array/Global.go Array/ArraySearchItem.go Array/ArraySearchItem_test.go
func BenchmarkArraySearchItem(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	var list []interface{}
	type item map[string]interface{}
	item1 := item{"age": 20, "name": "test1"}
	item2 := item{"age": 21, "name": "test2"}
	list = append(list, item1, item2, nil, "hello")
	var arr2 = map[string]interface{}{"age":21}
	for i:=0; i< t.N; i++ {
		_ = toolArray.ArraySearchItem(list, arr2)
	}
}