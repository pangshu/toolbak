package Date

import (
	"fmt"
	"testing"
	"time"
)

// 测试命令: go test -v -run TestStartOfYear Date/Global.go Date/StartOfYear.go Date/StartOfYear_test.go
func TestStartOfYear(t *testing.T) {
	var toolDate Date
	res1 := toolDate.StartOfYear(time.Now())
	fmt.Println(res1.Unix())
}

// go test -v -run TestStartOfYear -bench=BenchmarkStartOfYear -count=5 Date/Global.go Date/StartOfYear.go Date/StartOfYear_test.go
func BenchmarkStartOfYear(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_ = toolDate.StartOfYear(time.Now())
	}
}