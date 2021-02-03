package Date

import (
	"fmt"
	"testing"
	"time"
)

// 测试命令: go test -v -run TestFormat Date/Global.go Date/Format.go Date/Format_test.go
func TestFormat(t *testing.T) {
	var toolDate Date
	res1 := toolDate.Format("Y-m-d H:i:s", time.Now())
	fmt.Println(res1)
}

// go test -v -run TestFormat -bench=BenchmarkFormat -count=5 Date/Global.go Date/Format.go Date/Format_test.go
func BenchmarkFormat(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_ = toolDate.Format("Y-m-d H:i:s", time.Now())
	}
}