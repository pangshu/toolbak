package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsLink File/Global.go File/IsLink.go File/IsLink_test.go
func TestIsLink(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	res1 := toolFile.IsLink(form)
	fmt.Println(res1)
}

// go test -v -run TestIsLink -bench=BenchmarkIsLink -count=5 File/Global.go File/IsLink.go File/IsLink_test.go
func BenchmarkIsLink(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	for i:=0; i< t.N; i++ {
		_ = toolFile.IsLink(form)
	}
}