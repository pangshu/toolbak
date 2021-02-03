package Date

import (
	"fmt"
	"testing"
	"time"
)

// 测试命令: go test -v -run TestEndOfWeek Date/Global.go Date/EndOfWeek.go Date/EndOfWeek_test.go Date/StartOfWeek.go
func TestEndOfWeek(t *testing.T) {
	var toolDate Date
	res1 := toolDate.EndOfWeek(time.Now())
	fmt.Println(res1.Unix())
}

// go test -v -run TestEndOfWeek -bench=BenchmarkEndOfWeek -count=5 Date/Global.go Date/EndOfWeek.go Date/EndOfWeek_test.go Date/StartOfWeek.go
func BenchmarkEndOfWeek(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_ = toolDate.EndOfWeek(time.Now())
	}
}