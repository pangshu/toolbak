package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestDirSize File/Global.go File/DirSize.go File/DirSize_test.go File/IsDir.go
func TestDirSize(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/core/config/"
	res1 := toolFile.DirSize(form)
	fmt.Println(res1)
}

// go test -v -run TestDirSize -bench=BenchmarkDirSize -count=5 File/Global.go File/DirSize.go File/DirSize_test.go File/AbsPath.go File/IsDir.go
func BenchmarkDirSize(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v5/core/router"
	for i:=0; i< t.N; i++ {
		_ = toolFile.DirSize(form)
	}
}