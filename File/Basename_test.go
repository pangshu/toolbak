package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestBasename File/Global.go File/Basename.go File/Basename_test.go
func TestBasename(t *testing.T) {
	var toolFile File
	filename := "./test.go/aaa.go"
	res1 := toolFile.Basename(filename)
	fmt.Println(res1)
}

// go test -v -run TestBasename -bench=BenchmarkBasename -count=5 File/Global.go File/Basename.go File/Basename_test.go
func BenchmarkBasename(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	filename := "./test.go"
	for i:=0; i< t.N; i++ {
		_ = toolFile.Basename(filename)
	}
}