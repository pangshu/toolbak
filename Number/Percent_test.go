package Number

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestPercent Number/Global.go Number/Percent.go Number/Percent_test.go Number/ToFloat.go
func TestPercent(t *testing.T) {
	var toolNumber Number
	res1 := toolNumber.Percent(1, 1000)
	fmt.Println(res1)
}

// go test -v -run TestPercent -bench=BenchmarkPercent -count=5 Number/Global.go Number/Percent.go Number/Percent_test.go Number/ToFloat.go
func BenchmarkPercent(t *testing.B) {
	t.ResetTimer()
	var toolNumber Number
	for i:=0; i< t.N; i++ {
		_ = toolNumber.Percent(5, 10)
	}
}