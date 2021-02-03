package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestSumInt Number/Global.go Number/SumInt.go Number/SumInt_test.go
func TestSumInt(t *testing.T) {
	var toolNumber Number
	res1 := toolNumber.SumInt(5, 10)
	fmt.Println(res1)
}

// go test -v -run TestSumInt -bench=BenchmarkSumInt -count=5 Number/Global.go Number/SumInt.go Number/SumInt_test.go
func BenchmarkSumInt(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	for i:=0; i< t.N; i++ {
		_ = toolNumber.SumInt(5, 10)
	}
}