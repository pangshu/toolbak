package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestFileSize File/Global.go File/FileSize.go File/FileSize_test.go
func TestFileSize(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	res1 := toolFile.FileSize(form)
	fmt.Println(res1)
}

// go test -v -run TestFileSize -bench=BenchmarkFileSize -count=5 File/Global.go File/FileSize.go File/FileSize_test.go
func BenchmarkFileSize(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	for i:=0; i< t.N; i++ {
		_ = toolFile.FileSize(form)
	}
}