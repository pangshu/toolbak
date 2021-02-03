package Date

import (
	"fmt"
	"testing"
	"time"
)

// 测试命令: go test -v -run TestEndOfYear Date/Global.go Date/EndOfYear.go Date/EndOfYear_test.go Date/StartOfYear.go
func TestEndOfYear(t *testing.T) {
	var toolDate Date
	res1 := toolDate.EndOfYear(time.Now())
	fmt.Println(res1.Unix())
}

// go test -v -run TestEndOfYear -bench=BenchmarkEndOfYear -count=5 Date/Global.go Date/EndOfYear.go Date/EndOfYear_test.go Date/StartOfYear.go
func BenchmarkEndOfYear(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_ = toolDate.EndOfYear(time.Now())
	}
}