package Net

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestGetLocalMac Net/Global.go Net/GetLocalMac.go Net/GetLocalMac_test.go
func TestGetLocalMac(t *testing.T) {
	var toolNet Net
	res1 := toolNet.GetLocalMac()
	fmt.Println(res1)
}

// go test -v -run TestGetLocalMac -bench=BenchmarkGetLocalMac -count=5 Net/Global.go Net/GetLocalMac.go Net/GetLocalMac_test.go
func BenchmarkGetLocalMac(t *testing.B) {
	t.ResetTimer()
	var toolNet Net
	for i:=0; i< t.N; i++ {
		_ = toolNet.GetLocalMac()
	}
}