package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestPathInfo File/Global.go File/PathInfo.go File/PathInfo_test.go
func TestPathInfo(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/aaa/bbb/"
	res1 := toolFile.PathInfo(form, -1)
	fmt.Println(res1)
}

// go test -v -run TestPathInfo -bench=BenchmarkPathInfo -count=5 File/Global.go File/PathInfo.go File/PathInfo_test.go
func BenchmarkPathInfo(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	for i:=0; i< t.N; i++ {
		_ = toolFile.PathInfo(form, -1)
	}
}