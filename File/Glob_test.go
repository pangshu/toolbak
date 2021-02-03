package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestGlob File/Global.go File/Glob.go File/Glob_test.go
func TestGlob(t *testing.T) {
	var toolFile File
	form := "*test.go"
	res1,_ := toolFile.Glob(form)
	fmt.Println(res1)
}

// go test -v -run TestGlob -bench=BenchmarkGlob -count=5 File/Global.go File/Glob.go File/Glob_test.go
func BenchmarkGlob(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	for i:=0; i< t.N; i++ {
		_,_ = toolFile.Glob(form)
	}
}