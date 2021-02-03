package Date

import (
	"fmt"
	"testing"
	"time"
)

// 测试命令: go test -v -run TestHour Date/Global.go Date/Hour.go Date/Hour_test.go
func TestHour(t *testing.T) {
	var toolDate Date
	res1 := toolDate.Hour(time.Now())
	fmt.Println(res1)
}

// go test -v -run TestHour -bench=BenchmarkHour -count=5 Date/Global.go Date/Hour.go Date/Hour_test.go
func BenchmarkHour(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_ = toolDate.Hour(time.Now())
	}
}