package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestFormatDir File/Global.go File/FormatDir.go File/FormatDir_test.go
func TestFormatDir(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/"
	res1 := toolFile.FormatDir(form)
	fmt.Println(res1)
}

// go test -v -run TestFormatDir -bench=BenchmarkFormatDir -count=5 File/Global.go File/FormatDir.go File/FormatDir_test.go
func BenchmarkFormatDir(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	for i:=0; i< t.N; i++ {
		_ = toolFile.FormatDir(form)
	}
}