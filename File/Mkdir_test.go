package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestMkdir File/Global.go File/Mkdir.go File/Mkdir_test.go File/Ext.go File/IsExist.go
func TestMkdir(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/aaa/bbb/"
	res1 := toolFile.Mkdir(form, 0777)
	fmt.Println(res1)
}

// go test -v -run TestMkdir -bench=BenchmarkMkdir -count=5 File/Global.go File/Mkdir.go File/Mkdir_test.go
func BenchmarkMkdir(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	for i:=0; i< t.N; i++ {
		_ = toolFile.Mkdir(form, 0777)
	}
}