package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestFastCopy File/Global.go File/FastCopy.go File/FastCopy_test.go File/CopyLink.go File/CopyFile.go File/IsDir.go File/IsExist.go
func TestFastCopy(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v4"
	to := "/www/golang/src/v5"
	res1,_ := toolFile.FastCopy(form, to)
	fmt.Println(res1)
}

// go test -v -run TestFastCopy -bench=BenchmarkFastCopy -count=5 File/Global.go File/FastCopy.go File/FastCopy_test.go
func BenchmarkFastCopy(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	to := "/www/golang/src/v5"
	for i:=0; i< t.N; i++ {
		_,_ = toolFile.FastCopy(form, to)
	}
}