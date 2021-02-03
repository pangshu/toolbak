package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsWritable File/Global.go File/IsWritable.go File/IsWritable_test.go
func TestIsWritable(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	res1 := toolFile.IsWritable(form)
	fmt.Println(res1)
}

// go test -v -run TestIsWritable -bench=BenchmarkIsWritable -count=5 File/Global.go File/IsWritable.go File/IsWritable_test.go
func BenchmarkIsWritable(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	for i:=0; i< t.N; i++ {
		_ = toolFile.IsWritable(form)
	}
}