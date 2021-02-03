package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsBinary File/Global.go File/IsBinary.go File/IsBinary_test.go File/ReadFile.go
func TestIsBinary(t *testing.T) {
	var toolFile File
	form := "./Global.go"
	res1 := toolFile.IsBinary(form)
	fmt.Println(res1)
}

// go test -v -run TestIsBinary -bench=BenchmarkIsBinary -count=5 File/Global.go File/IsBinary.go File/IsBinary_test.go
func BenchmarkIsBinary(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	for i:=0; i< t.N; i++ {
		_ = toolFile.IsBinary(form)
	}
}