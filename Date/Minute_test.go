package Date

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestMinute Date/Global.go Date/Minute.go Date/Minute_test.go
func TestMinute(t *testing.T) {
	var toolDate Date
	res1 := toolDate.Minute()
	fmt.Println(res1)
}

// go test -v -run TestMinute -bench=BenchmarkMinute -count=5 Date/Global.go Date/Minute.go Date/Minute_test.go
func BenchmarkMinute(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_ = toolDate.Minute()
	}
}