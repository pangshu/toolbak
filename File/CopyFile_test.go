package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestCopyFile File/Global.go File/CopyFile.go File/CopyFile_test.go File/IsDir.go File/IsExist.go
func TestCopyFile(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	to := "/www/golang/src/v5/aaa.go"
	res1,_ := toolFile.CopyFile(form, to)
	fmt.Println(res1)
}

// go test -v -run TestCopyFile -bench=BenchmarkCopyFile -count=5 File/Global.go File/CopyFile.go File/CopyFile_test.go
func BenchmarkCopyFile(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	to := "/www/golang/src/v5"
	for i:=0; i< t.N; i++ {
		_,_ = toolFile.CopyFile(form, to)
	}
}