package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsReadable File/Global.go File/IsReadable.go File/IsReadable_test.go
func TestIsReadable(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	res1 := toolFile.IsReadable(form)
	fmt.Println(res1)
}

// go test -v -run TestIsReadable -bench=BenchmarkIsReadable -count=5 File/Global.go File/IsReadable.go File/IsReadable_test.go
func BenchmarkIsReadable(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	for i:=0; i< t.N; i++ {
		_ = toolFile.IsReadable(form)
	}
}