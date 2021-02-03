package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsExecutable File/Global.go File/IsExecutable.go File/IsExecutable_test.go
func TestIsExecutable(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	res1 := toolFile.IsExecutable(form)
	fmt.Println(res1)
}

// go test -v -run TestIsExecutable -bench=BenchmarkIsExecutable -count=5 File/Global.go File/IsExecutable.go File/IsExecutable_test.go
func BenchmarkIsExecutable(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	for i:=0; i< t.N; i++ {
		_ = toolFile.IsExecutable(form)
	}
}