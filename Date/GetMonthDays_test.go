package Date

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestGetMonthDays Date/Global.go Date/GetMonthDays.go Date/GetMonthDays_test.go
func TestGetMonthDays(t *testing.T) {
	var toolDate Date
	res1 := toolDate.GetMonthDays(2)
	fmt.Println(res1)
}

// go test -v -run TestGetMonthDays -bench=BenchmarkGetMonthDays -count=5 Date/Global.go Date/GetMonthDays.go Date/GetMonthDays_test.go
func BenchmarkGetMonthDays(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_ = toolDate.GetMonthDays(2)
	}
}