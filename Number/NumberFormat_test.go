package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestNumberFormat Number/Global.go Number/NumberFormat.go Number/NumberFormat_test.go Number/SumFloat64.go
func TestNumberFormat(t *testing.T) {
	var toolNumber Number
	sum := toolNumber.SumFloat64(0.0, 1.1, -2.2, 3.30, 5.55)
	res1 := toolNumber.NumberFormat(sum, 2, ".", ">")
	fmt.Println(res1)
}

// go test -v -run TestNumberFormat -bench=BenchmarkNumberFormat -count=5 Number/Global.go Number/NumberFormat.go Number/NumberFormat_test.go
func BenchmarkNumberFormat(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	sum := toolNumber.SumFloat64(0.0, 1.1, -2.2, 3.30, 5.55)
	for i:=0; i< t.N; i++ {
		_ = toolNumber.NumberFormat(sum, 2, ".", "")
	}
}