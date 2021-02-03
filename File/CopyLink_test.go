package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestCopyLink File/Global.go File/CopyLink.go File/CopyLink_test.go File/IsDir.go File/IsExist.go
func TestCopyLink(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/bb.go"
	to := "/www/golang/src/v5/bbb.go"
	res1 := toolFile.CopyLink(form, to)
	fmt.Println(res1)
}

// go test -v -run TestCopyLink -bench=BenchmarkCopyLink -count=5 File/Global.go File/CopyLink.go File/CopyLink_test.go
func BenchmarkCopyLink(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	to := "/www/golang/src/v5"
	for i:=0; i< t.N; i++ {
		_ = toolFile.CopyLink(form, to)
	}
}