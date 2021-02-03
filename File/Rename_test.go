package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestRename File/Global.go File/Rename.go File/Rename_test.go
func TestRename(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main111.go"
	res1 := toolFile.Rename(form, "main222.go")
	fmt.Println(res1)
}

// go test -v -run TestRename -bench=BenchmarkRename -count=5 File/Global.go File/Rename.go File/Rename_test.go
func BenchmarkRename(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	for i:=0; i< t.N; i++ {
		_ = toolFile.Rename(form, "main222.go")
	}
}