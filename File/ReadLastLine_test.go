package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestReadLastLine File/Global.go File/ReadLastLine.go File/ReadLastLine_test.go
func TestReadLastLine(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/aaa.go"
	res1 := toolFile.ReadLastLine(form)
	fmt.Println(res1)
}

// go test -v -run TestReadLastLine -bench=BenchmarkReadLastLine -count=5 File/Global.go File/ReadLastLine.go File/ReadLastLine_test.go
func BenchmarkReadLastLine(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	for i:=0; i< t.N; i++ {
		_ = toolFile.ReadLastLine(form)
	}
}