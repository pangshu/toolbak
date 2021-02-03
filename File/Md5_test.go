package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestMd5 File/Global.go File/Md5.go File/Md5_test.go File/Ext.go File/IsExist.go
func TestMd5(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	res1,_ := toolFile.Md5(form, 32)
	fmt.Println(res1)
}

// go test -v -run TestMd5 -bench=BenchmarkMd5 -count=5 File/Global.go File/Md5.go File/Md5_test.go
func BenchmarkMd5(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	for i:=0; i< t.N; i++ {
		_,_ = toolFile.Md5(form)
	}
}