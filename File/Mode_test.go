package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestMode File/Global.go File/Mode.go File/Mode_test.go
func TestMode(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/aaa/bbb/"
	res1,_ := toolFile.Mode(form)
	fmt.Println(res1)
}

// go test -v -run TestMode -bench=BenchmarkMode -count=5 File/Global.go File/Mode.go File/Mode_test.go
func BenchmarkMode(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	for i:=0; i< t.N; i++ {
		_,_ = toolFile.Mode(form)
	}
}