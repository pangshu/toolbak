package Date

import (
	"fmt"
	"testing"
	"time"
)

// 测试命令: go test -v -run TestStartOfMonth Date/Global.go Date/StartOfMonth.go Date/StartOfMonth_test.go
func TestStartOfMonth(t *testing.T) {
	var toolDate Date
	res1 := toolDate.StartOfMonth(time.Now())
	fmt.Println(res1.Unix())
}

// go test -v -run TestStartOfMonth -bench=BenchmarkStartOfMonth -count=5 Date/Global.go Date/StartOfMonth.go Date/StartOfMonth_test.go
func BenchmarkStartOfMonth(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_ = toolDate.StartOfMonth(time.Now())
	}
}