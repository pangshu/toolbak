package Date

import (
	"fmt"
	"testing"
	"time"
)

// 测试命令: go test -v -run TestStartOfDay Date/Global.go Date/StartOfDay.go Date/StartOfDay_test.go
func TestStartOfDay(t *testing.T) {
	var toolDate Date
	res1 := toolDate.StartOfDay(time.Now())
	fmt.Println(res1.Unix())
}

// go test -v -run TestStartOfDay -bench=BenchmarkStartOfDay -count=5 Date/Global.go Date/StartOfDay.go Date/StartOfDay_test.go
func BenchmarkStartOfDay(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_ = toolDate.StartOfDay(time.Now())
	}
}