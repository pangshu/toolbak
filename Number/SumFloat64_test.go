package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestSumFloat64 Number/Global.go Number/SumFloat64.go Number/SumFloat64_test.go
func TestSumFloat64(t *testing.T) {
	var toolNumber Number
	res1 := toolNumber.SumFloat64(5, 10)
	fmt.Println(res1)
}

// go test -v -run TestSumFloat64 -bench=BenchmarkSumFloat64 -count=5 Number/Global.go Number/SumFloat64.go Number/SumFloat64_test.go
func BenchmarkSumFloat64(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	for i:=0; i< t.N; i++ {
		_ = toolNumber.SumFloat64(5, 10)
	}
}