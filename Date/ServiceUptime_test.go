package Date

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestServiceUptime Date/Global.go Date/ServiceUptime.go Date/ServiceUptime_test.go
func TestServiceUptime(t *testing.T) {
	var toolDate Date
	res1 := toolDate.ServiceUptime()
	fmt.Println(res1)
}

// go test -v -run TestServiceUptime -bench=BenchmarkServiceUptime -count=5 Date/Global.go Date/ServiceUptime.go Date/ServiceUptime_test.go
func BenchmarkServiceUptime(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_ = toolDate.ServiceUptime()
	}
}