package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestHomePath Os/Global.go Os/HomePath.go Os/HomePath_test.go Os/IsWindows.go
func TestHomePath(t *testing.T) {
	var toolOs Os
	res1,_ := toolOs.HomePath()
	fmt.Println(res1)
}

// go test -v -run TestHomePath -bench=BenchmarkHomePath -count=5 Os/Global.go Os/HomePath.go Os/HomePath_test.go Os/IsWindows.go
func BenchmarkHomePath(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	for i:=0; i< t.N; i++ {
		_,_ = toolOs.HomePath()
	}
}