package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestGetFileMode File/Global.go File/GetFileMode.go File/GetFileMode_test.go
func TestGetFileMode(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	res1,_ := toolFile.GetFileMode(form)
	fmt.Println(res1)
}

// go test -v -run TestGetFileMode -bench=BenchmarkGetFileMode -count=5 File/Global.go File/GetFileMode.go File/GetFileMode_test.go
func BenchmarkGetFileMode(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	for i:=0; i< t.N; i++ {
		_,_ = toolFile.GetFileMode(form)
	}
}