package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestSetEnv Os/Global.go Os/SetEnv.go Os/SetEnv_test.go
func TestSetEnv(t *testing.T) {
	var toolOs Os
	res1 := toolOs.SetEnv("aaa", "bbb")
	fmt.Println(res1)
}

// go test -v -run TestSetEnv -bench=BenchmarkSetEnv -count=5 Os/Global.go Os/SetEnv.go Os/SetEnv_test.go
func BenchmarkSetEnv(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	for i:=0; i< t.N; i++ {
		_ = toolOs.SetEnv("aaa", "bbb")
	}
}