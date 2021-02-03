package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestRealPath File/Global.go File/RealPath.go File/RealPath_test.go
func TestRealPath(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/aaaaa.go"
	res1 := toolFile.RealPath(form)
	fmt.Println(res1)
}

// go test -v -run TestRealPath -bench=BenchmarkRealPath -count=5 File/Global.go File/RealPath.go File/RealPath_test.go
func BenchmarkRealPath(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	for i:=0; i< t.N; i++ {
		_ = toolFile.RealPath(form)
	}
}