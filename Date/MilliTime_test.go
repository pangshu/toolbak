package Date

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestMilliTime Date/Global.go Date/MilliTime.go Date/MilliTime_test.go
func TestMilliTime(t *testing.T) {
	var toolDate Date
	res1 := toolDate.MilliTime()
	fmt.Println(res1)
}

// go test -v -run TestMilliTime -bench=BenchmarkMilliTime -count=5 Date/Global.go Date/MilliTime.go Date/MilliTime_test.go
func BenchmarkMilliTime(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_ = toolDate.MilliTime()
	}
}