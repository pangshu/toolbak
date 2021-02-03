package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestGetPidByPort Os/Global.go Os/GetPidByPort.go Os/GetPidByPort_test.go
func TestGetPidByPort(t *testing.T) {
	var toolOs Os
	res1 := toolOs.GetPidByPort(22)
	fmt.Println(res1)
}

// go test -v -run TestGetPidByPort -bench=BenchmarkGetPidByPort -count=5 Os/Global.go Os/GetPidByPort.go Os/GetPidByPort_test.go
func BenchmarkGetPidByPort(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	for i:=0; i< t.N; i++ {
		_ = toolOs.GetPidByPort(22)
	}
}