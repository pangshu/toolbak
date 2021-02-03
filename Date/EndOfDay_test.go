package Date

import (
	"fmt"
	"testing"
	"time"
)

// 测试命令: go test -v -run TestEndOfDay Date/Global.go Date/EndOfDay.go Date/EndOfDay_test.go
func TestEndOfDay(t *testing.T) {
	var toolDate Date
	res1 := toolDate.EndOfDay(time.Now())
	fmt.Println(res1.Unix())
}

// go test -v -run TestEndOfDay -bench=BenchmarkEndOfDay -count=5 Date/Global.go Date/EndOfDay.go Date/EndOfDay_test.go
func BenchmarkEndOfDay(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_ = toolDate.EndOfDay(time.Now())
	}
}