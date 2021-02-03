package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestRemove File/Global.go File/Remove.go File/Remove_test.go
func TestRemove(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/aaa.go"
	res1 := toolFile.Remove(form)
	fmt.Println(res1)
}

// go test -v -run TestRemove -bench=BenchmarkRemove -count=5 File/Global.go File/Remove.go File/Remove_test.go
func BenchmarkRemove(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	for i:=0; i< t.N; i++ {
		_ = toolFile.Remove(form)
	}
}