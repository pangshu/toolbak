package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestGetProcessPathByPid Os/Global.go Os/GetProcessPathByPid.go Os/GetProcessPathByPid_test.go
func TestGetProcessPathByPid(t *testing.T) {
	var toolOs Os
	res1 := toolOs.GetProcessPathByPid(1461)
	fmt.Println(res1)
}

// go test -v -run TestGetProcessPathByPid -bench=BenchmarkGetProcessPathByPid -count=5 Os/Global.go Os/GetProcessPathByPid.go Os/GetProcessPathByPid_test.go
func BenchmarkGetProcessPathByPid(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	for i:=0; i< t.N; i++ {
		_ = toolOs.GetProcessPathByPid(1461)
	}
}