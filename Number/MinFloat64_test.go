package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestMinFloat64 Number/Global.go Number/MinFloat64.go Number/MinFloat64_test.go
func TestMinFloat64(t *testing.T) {
	var toolNumber Number
	num := []float64{-4, 0, 3, 9}
	res1 := toolNumber.MinFloat64(num...)
	fmt.Println(res1)
}

// go test -v -run TestMinFloat64 -bench=BenchmarkMinFloat64 -count=5 Number/Global.go Number/MinFloat64.go Number/MinFloat64_test.go
func BenchmarkMinFloat64(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	num := []float64{-4, 0, 3, 9}
	for i:=0; i< t.N; i++ {
		_ = toolNumber.MinFloat64(num...)
	}
}