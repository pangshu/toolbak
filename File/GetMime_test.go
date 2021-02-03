package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestGetMime File/Global.go File/GetMime.go File/GetMime_test.go
func TestGetMime(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	res1 := toolFile.GetMime(form, false)
	fmt.Println(res1)
}

// go test -v -run TestGetMime -bench=BenchmarkGetMime -count=5 File/Global.go File/GetMime.go File/GetMime_test.go
func BenchmarkGetMime(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	for i:=0; i< t.N; i++ {
		_ = toolFile.GetMime(form, false)
	}
}