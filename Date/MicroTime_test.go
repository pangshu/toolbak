package Date

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestMicroTime Date/Global.go Date/MicroTime.go Date/MicroTime_test.go
func TestMicroTime(t *testing.T) {
	var toolDate Date
	res1 := toolDate.MicroTime()
	fmt.Println(res1)
}

// go test -v -run TestMicroTime -bench=BenchmarkMicroTime -count=5 Date/Global.go Date/MicroTime.go Date/MicroTime_test.go
func BenchmarkMicroTime(t *testing.B) {
	t.ResetTimer()
	var toolDate Date
	for i:=0; i< t.N; i++ {
		_ = toolDate.MicroTime()
	}
}