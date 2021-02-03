package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestByteFormat Os/Global.go Os/ByteFormat.go Os/ByteFormat_test.go
func TestByteFormat(t *testing.T) {
	var toolOs Os
	var num float64 = 15340331
	res1 := toolOs.ByteFormat(num, 2, " ")
	fmt.Println(res1)
}

// go test -v -run TestByteFormat -bench=BenchmarkByteFormat -count=5 Os/Global.go Os/ByteFormat.go Os/ByteFormat_test.go
func BenchmarkByteFormat(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	var num float64 = 15340331
	for i:=0; i< t.N; i++ {
		_ = toolOs.ByteFormat(num, 2, " ")
	}
}