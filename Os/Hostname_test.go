package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestHostname Os/Global.go Os/Hostname.go Os/Hostname_test.go
func TestHostname(t *testing.T) {
	var toolOs Os
	res1,_ := toolOs.Hostname()
	fmt.Println(res1)
}

// go test -v -run TestHostname -bench=BenchmarkHostname -count=5 Os/Global.go Os/Hostname.go Os/Hostname_test.go
func BenchmarkHostname(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	for i:=0; i< t.N; i++ {
		_,_ = toolOs.Hostname()
	}
}