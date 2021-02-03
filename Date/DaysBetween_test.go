package Date

import (
	"fmt"
	"testing"
	"time"
)

// 测试命令: go test -v -run TestDaysBetween Date/Global.go Date/DaysBetween.go Date/DaysBetween_test.go
func TestDaysBetween(t *testing.T) {
	var toolDate Date
	res1 := toolDate.DaysBetween(time.Unix(1610236800,0), time.Now())
	fmt.Println(res1)
}

// go test -v -run TestDaysBetween -bench=BenchmarkDaysBetween -count=5 Date/Global.go Date/DaysBetween.go Date/DaysBetween_test.go
func BenchmarkDaysBetween(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_ = toolDate.DaysBetween(time.Unix(1610236800,0), time.Now())
	}
}