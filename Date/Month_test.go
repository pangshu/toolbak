package Date

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestMonth Date/Global.go Date/Month.go Date/Month_test.go
func TestMonth(t *testing.T) {
	var toolDate Date
	res1 := toolDate.Month()
	fmt.Println(res1)
}

// go test -v -run TestMonth -bench=BenchmarkMonth -count=5 Date/Global.go Date/Month.go Date/Month_test.go
func BenchmarkMonth(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_ = toolDate.Month()
	}
}