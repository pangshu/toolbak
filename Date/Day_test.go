package Date

import (
	"fmt"
	"testing"
	"time"
)

// 测试命令: go test -v -run TestDay Date/Global.go Date/Day.go Date/Day_test.go
func TestDay(t *testing.T) {
	var toolDate Date
	res1 := toolDate.Day(time.Now())
	fmt.Println(res1)
}

// go test -v -run TestDay -bench=BenchmarkDay -count=5 Date/Global.go Date/Day.go Date/Day_test.go
func BenchmarkDay(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_ = toolDate.Day(time.Now())
	}
}