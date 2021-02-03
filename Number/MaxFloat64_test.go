package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestMaxFloat64 Number/Global.go Number/MaxFloat64.go Number/MaxFloat64_test.go
func TestMaxFloat64(t *testing.T) {
	var toolNumber Number
	num := []float64{-4, 0, 3, 9}
	res1 := toolNumber.MaxFloat64(num...)
	fmt.Println(res1)
}

// go test -v -run TestMaxFloat64 -bench=BenchmarkMaxFloat64 -count=5 Number/Global.go Number/MaxFloat64.go Number/MaxFloat64_test.go
func BenchmarkMaxFloat64(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	num := []float64{-4, 0, 3, 9}
	for i:=0; i< t.N; i++ {
		_ = toolNumber.MaxFloat64(num...)
	}
}