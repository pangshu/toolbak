package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestFormatPath File/Global.go File/FormatPath.go File/FormatPath_test.go File/FormatDir.go
func TestFormatPath(t *testing.T) {
	var toolFile File
	form := `/usr|///tmp:\\\123/\abc<|\hello>\/%world?\\how$\\are\@#test.png`
	res1 := toolFile.FormatPath(form)
	fmt.Println(res1)
}

// go test -v -run TestFormatPath -bench=BenchmarkFormatPath -count=5 File/Global.go File/FormatPath.go File/FormatPath_test.go
func BenchmarkFormatPath(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	for i:=0; i< t.N; i++ {
		_ = toolFile.FormatPath(form)
	}
}