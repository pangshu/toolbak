package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestWriteFile File/Global.go File/WriteFile.go File/WriteFile_test.go File/IsDir.go
func TestWriteFile(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main333.go"
	res1 := toolFile.WriteFile(form, "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	fmt.Println(res1)
}

// go test -v -run TestWriteFile -bench=BenchmarkWriteFile -count=5 File/Global.go File/WriteFile.go File/WriteFile_test.go
func BenchmarkWriteFile(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	for i:=0; i< t.N; i++ {
		_ = toolFile.WriteFile(form, "AAAAAAAAAAAAAAAAAAAAa")
	}
}