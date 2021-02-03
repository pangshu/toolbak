package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestMinInt Number/Global.go Number/MinInt.go Number/MinInt_test.go
func TestMinInt(t *testing.T) {
	var toolNumber Number
	num := []int{-4, 0, 3, 9}
	res1 := toolNumber.MinInt(num...)
	fmt.Println(res1)
}

// go test -v -run TestMinInt -bench=BenchmarkMinInt -count=5 Number/Global.go Number/MinInt.go Number/MinInt_test.go
func BenchmarkMinInt(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	num := []int{-4, 0, 3, 9}
	for i:=0; i< t.N; i++ {
		_ = toolNumber.MinInt(num...)
	}
}