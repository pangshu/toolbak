package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestDelDir File/Global.go File/DelDir.go File/DelDir_test.go File/AbsPath.go File/IsDir.go
func TestDelDir(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/core/router"
	res1 := toolFile.DelDir(form)
	fmt.Println(res1)
}

// go test -v -run TestDelDir -bench=BenchmarkDelDir -count=5 File/Global.go File/DelDir.go File/DelDir_test.go File/AbsPath.go File/IsDir.go
func BenchmarkDelDir(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v5/core/router"
	for i:=0; i< t.N; i++ {
		_ = toolFile.DelDir(form)
	}
}