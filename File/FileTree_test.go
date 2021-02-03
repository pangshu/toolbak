package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestFileTree File/Global.go File/FileTree.go File/FileTree_test.go File/IsFile.go File/IsDir.go File/IsExist.go
func TestFileTree(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/"
	res1 := toolFile.FileTree(form, 0 ,true)
	fmt.Println(res1)
}

// go test -v -run TestFileTree -bench=BenchmarkFileTree -count=5 File/Global.go File/FileTree.go File/FileTree_test.go
func BenchmarkFileTree(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	for i:=0; i< t.N; i++ {
		_ = toolFile.FileTree(form, 0 ,true)
	}
}