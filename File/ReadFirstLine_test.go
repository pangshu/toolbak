package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestReadFirstLine File/Global.go File/ReadFirstLine.go File/ReadFirstLine_test.go
func TestReadFirstLine(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	res1 := toolFile.ReadFirstLine(form)
	fmt.Println(string(res1))
}

// go test -v -run TestReadFirstLine -bench=BenchmarkReadFirstLine -count=5 File/Global.go File/ReadFirstLine.go File/ReadFirstLine_test.go
func BenchmarkReadFirstLine(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	for i:=0; i< t.N; i++ {
		_ = toolFile.ReadFirstLine(form)
	}
}