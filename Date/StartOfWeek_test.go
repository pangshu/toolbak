package Date

import (
	"fmt"
	"testing"
	"time"
)

// 测试命令: go test -v -run TestStartOfWeek Date/Global.go Date/StartOfWeek.go Date/StartOfWeek_test.go
func TestStartOfWeek(t *testing.T) {
	var toolDate Date
	res1 := toolDate.StartOfWeek(time.Now())
	fmt.Println(res1.Unix())
}

// go test -v -run TestStartOfWeek -bench=BenchmarkStartOfWeek -count=5 Date/Global.go Date/StartOfWeek.go Date/StartOfWeek_test.go
func BenchmarkStartOfWeek(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_ = toolDate.StartOfWeek(time.Now())
	}
}