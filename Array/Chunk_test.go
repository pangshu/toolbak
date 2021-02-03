package Array

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestChunk Array/Global.go Array/Chunk.go Array/Chunk_test.go
func TestChunk(t *testing.T) {
	var toolArray Array
	var arr1 = []interface{}{1,2,3,4,5,6,7,8,9,1}
	res1 := toolArray.Chunk(arr1,2)
	fmt.Println(res1)
}

// go test -v -run TestChunk -bench=BenchmarkChunk -count=5 Array/Global.go Array/Chunk.go Array/Chunk_test.go
func BenchmarkChunk(t *testing.B) {
	t.ResetTimer()
	var toolArray Array
	var arr1 = []interface{}{1,2,3,4,5,6,7,8,9,1}
	for i:=0; i< t.N; i++ {
		_ = toolArray.Chunk(arr1,2)
	}
}