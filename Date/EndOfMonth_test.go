package Date

import (
	"fmt"
	"testing"
	"time"
)

// 测试命令: go test -v -run TestEndOfMonth Date/Global.go Date/EndOfMonth.go Date/EndOfMonth_test.go Date/StartOfMonth.go
func TestEndOfMonth(t *testing.T) {
	var toolDate Date
	res1 := toolDate.EndOfMonth(time.Now())
	fmt.Println(res1.Unix())
}

// go test -v -run TestEndOfMonth -bench=BenchmarkEndOfMonth -count=5 Date/Global.go Date/EndOfMonth.go Date/EndOfMonth_test.go Date/StartOfMonth.go
func BenchmarkEndOfMonth(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_ = toolDate.EndOfMonth(time.Now())
	}
}