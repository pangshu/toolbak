package Os

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestChmod Os/Global.go Os/Chmod.go Os/Chmod_test.go
func TestChmod(t *testing.T) {
	var toolOs Os
	file := "/www/golang/src/toolbak/test.go"
	res1 := toolOs.Chmod(file, 0777)
	fmt.Println(res1)
}

// go test -v -run TestChmod -bench=BenchmarkChmod -count=5 Os/Global.go Os/Chmod.go Os/Chmod_test.go
func BenchmarkChmod(t *testing.B) {
	t.ResetTimer()
	var toolOs Os
	file := "/www/golang/src/toolbak/test.go"
	for i:=0; i< t.N; i++ {
		_ = toolOs.Chmod(file, 0777)
	}
}