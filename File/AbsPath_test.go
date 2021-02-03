package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestAbsPath File/Global.go File/AbsPath.go File/AbsPath_test.go
func TestAbsPath(t *testing.T) {
	var toolFile File
	str := "./test.go"
	res1 := toolFile.AbsPath(str)
	fmt.Println(res1)
}

// go test -v -run TestAbsPath -bench=BenchmarkAbsPath -count=5 File/Global.go File/AbsPath.go File/AbsPath_test.go
func BenchmarkAbsPath(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	str := "./test.go"
	for i:=0; i< t.N; i++ {
		_ = toolFile.AbsPath(str)
	}
}