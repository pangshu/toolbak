package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsDir File/Global.go File/IsDir.go File/IsDir_test.go
func TestIsDir(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5"
	res1 := toolFile.IsDir(form)
	fmt.Println(res1)
}

// go test -v -run TestIsDir -bench=BenchmarkIsDir -count=5 File/Global.go File/IsDir.go File/IsDir_test.go
func BenchmarkIsDir(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	for i:=0; i< t.N; i++ {
		_ = toolFile.IsDir(form)
	}
}