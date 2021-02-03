package Net

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsMacAddr Net/Global.go Net/IsMacAddr.go Net/IsMacAddr_test.go
func TestIsMacAddr(t *testing.T) {
	var toolNet Net
	res1 := toolNet.IsMacAddr("2001:A304:6101:1::E0:F726:4E58")
	fmt.Println(res1)
}

// go test -v -run TestIsMacAddr -bench=BenchmarkIsMacAddr -count=5 Net/Global.go Net/IsMacAddr.go Net/IsMacAddr_test.go
func BenchmarkIsMacAddr(t *testing.B) {
	t.ResetTimer()
	var toolNet Net
	for i:=0; i< t.N; i++ {
		_ = toolNet.IsMacAddr("2001:A304:6101:1::E0:F726:4E58")
	}
}