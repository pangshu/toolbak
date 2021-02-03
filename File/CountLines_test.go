package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestCountLines File/Global.go File/CountLines.go File/CountLines_test.go
func TestCountLines(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	res1,_ := toolFile.CountLines(form, 8)
	fmt.Println(res1)
}

// go test -v -run TestCountLines -bench=BenchmarkCountLines -count=5 File/Global.go File/CountLines.go File/CountLines_test.go
func BenchmarkCountLines(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	for i:=0; i< t.N; i++ {
		_,_ = toolFile.CountLines(form, 8)
	}
}