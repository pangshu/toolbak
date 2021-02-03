package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestGetEnv Os/Global.go Os/GetEnv.go Os/GetEnv_test.go
func TestGetEnv(t *testing.T) {
	var toolOs Os
	pathName := "GOROOT"
	res1 := toolOs.GetEnv(pathName)
	fmt.Println(res1)
}

// go test -v -run TestGetEnv -bench=BenchmarkGetEnv -count=5 Os/Global.go Os/GetEnv.go Os/GetEnv_test.go
func BenchmarkGetEnv(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	pathName := "GOROOT"
	for i:=0; i< t.N; i++ {
		_ = toolOs.GetEnv(pathName)
	}
}