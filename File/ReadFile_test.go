package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestReadFile File/Global.go File/ReadFile.go File/ReadFile_test.go
func TestReadFile(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	res1,_ := toolFile.ReadFile(form)
	fmt.Println(string(res1))
}

// go test -v -run TestReadFile -bench=BenchmarkReadFile -count=5 File/Global.go File/ReadFile.go File/ReadFile_test.go
func BenchmarkReadFile(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	for i:=0; i< t.N; i++ {
		_,_ = toolFile.ReadFile(form)
	}
}