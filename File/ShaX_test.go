package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestShaX File/Global.go File/ShaX.go File/ShaX_test.go
func TestShaX(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main222.go"
	res1,_ := toolFile.ShaX(form, 1)
	fmt.Println(res1)
}

// go test -v -run TestShaX -bench=BenchmarkShaX -count=5 File/Global.go File/ShaX.go File/ShaX_test.go
func BenchmarkShaX(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	for i:=0; i< t.N; i++ {
		_,_ = toolFile.ShaX(form,1)
	}
}