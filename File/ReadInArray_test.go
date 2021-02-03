package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestReadInArray File/Global.go File/ReadInArray.go File/ReadInArray_test.go
func TestReadInArray(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	res1,_ := toolFile.ReadInArray(form)
	fmt.Println(res1)
}

// go test -v -run TestReadInArray -bench=BenchmarkReadInArray -count=5 File/Global.go File/ReadInArray.go File/ReadInArray_test.go
func BenchmarkReadInArray(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	for i:=0; i< t.N; i++ {
		_,_ = toolFile.ReadInArray(form)
	}
}