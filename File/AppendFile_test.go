package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestAppendFile File/Global.go File/AppendFile.go File/AppendFile_test.go File/Mode.go
func TestAppendFile(t *testing.T) {
	var toolFile File
	filename := "./test.go"
	str := "测试测试"
	res1 := toolFile.AppendFile(filename, str)
	fmt.Println(res1)
}

// go test -v -run TestAppendFile -bench=BenchmarkAppendFile -count=5 File/Global.go File/AppendFile.go File/AppendFile_test.go File/Mode.go
func BenchmarkAppendFile(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	filename := "./test.go"
	str := "测试测试"
	for i:=0; i< t.N; i++ {
		_ = toolFile.AppendFile(filename, str)
	}
}