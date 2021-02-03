package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestGetModTime File/Global.go File/GetModTime.go File/GetModTime_test.go
func TestGetModTime(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	res1 := toolFile.GetModTime(form)
	fmt.Println(res1)
}

// go test -v -run TestGetModTime -bench=BenchmarkGetModTime -count=5 File/Global.go File/GetModTime.go File/GetModTime_test.go
func BenchmarkGetModTime(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	for i:=0; i< t.N; i++ {
		_ = toolFile.GetModTime(form)
	}
}