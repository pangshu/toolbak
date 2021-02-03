package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestCopyDir File/Global.go File/CopyDir.go File/CopyDir_test.go File/CopyLink.go File/CopyFile.go File/IsDir.go File/IsExist.go
func TestCopyDir(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v4"
	to := "/www/golang/src/v5"
	res1,_ := toolFile.CopyDir(form, to)
	fmt.Println(res1)
}

// go test -v -run TestCopyDir -bench=BenchmarkCopyDir -count=5 File/Global.go File/CopyDir.go File/CopyDir_test.go
func BenchmarkCopyDir(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	to := "/www/golang/src/v5"
	for i:=0; i< t.N; i++ {
		_,_ = toolFile.CopyDir(form, to)
	}
}