package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestSafeFileName File/Global.go File/SafeFileName.go File/SafeFileName_test.go
func TestSafeFileName(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main222.go"
	res1 := toolFile.SafeFileName(form)
	fmt.Println(res1)
}

// go test -v -run TestSafeFileName -bench=BenchmarkSafeFileName -count=5 File/Global.go File/SafeFileName.go File/SafeFileName_test.go
func BenchmarkSafeFileName(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	for i:=0; i< t.N; i++ {
		_ = toolFile.SafeFileName(form)
	}
}