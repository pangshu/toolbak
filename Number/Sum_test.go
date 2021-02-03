package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestSum Number/Global.go Number/Sum.go Number/Sum_test.go Number/ToFloat.go
func TestSum(t *testing.T) {
	var toolNumber Number
	res1 := toolNumber.Sum(5, 10)
	fmt.Println(res1)
}

// go test -v -run TestSum -bench=BenchmarkSum -count=5 Number/Global.go Number/Sum.go Number/Sum_test.go Number/ToFloat.go
func BenchmarkSum(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	for i:=0; i< t.N; i++ {
		_ = toolNumber.Sum(5, 10)
	}
}