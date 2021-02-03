package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestAverageFloat64 Number/Global.go Number/AverageFloat64.go Number/AverageFloat64_test.go Number/SumFloat64.go
func TestAverageFloat64(t *testing.T) {
	var toolNumber Number
	num := []float64{1,2.2,3.3,4,5,6,7,8,9,10}
	res1 := toolNumber.AverageFloat64(num...)
	fmt.Println(res1)
}

// go test -v -run TestAverageFloat64 -bench=BenchmarkAverageFloat64 -count=5 Number/Global.go Number/AverageFloat64.go Number/AverageFloat64_test.go Number/SumFloat64.go
func BenchmarkAverageFloat64(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	num := []float64{1,2,3,4,5,6,7,8,9,10}
	for i:=0; i< t.N; i++ {
		_ = toolNumber.AverageFloat64(num...)
	}
}