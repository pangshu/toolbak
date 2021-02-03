package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsFile File/Global.go File/IsFile.go File/IsFile_test.go File/IsExist.go
func TestIsFile(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	res1 := toolFile.IsFile(form)
	fmt.Println(res1)
}

// go test -v -run TestIsFile -bench=BenchmarkIsFile -count=5 File/Global.go File/IsFile.go File/IsFile_test.go
func BenchmarkIsFile(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	for i:=0; i< t.N; i++ {
		_ = toolFile.IsFile(form)
	}
}