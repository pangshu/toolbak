package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestDirname File/Global.go File/Dirname.go File/Dirname_test.go File/IsDir.go
func TestDirname(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/core/config/db.sql"
	res1 := toolFile.Dirname(form)
	fmt.Println(res1)
}

// go test -v -run TestDirname -bench=BenchmarkDirname -count=5 File/Global.go File/Dirname.go File/Dirname_test.go File/AbsPath.go File/IsDir.go
func BenchmarkDirname(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v5/core/router"
	for i:=0; i< t.N; i++ {
		_ = toolFile.Dirname(form)
	}
}