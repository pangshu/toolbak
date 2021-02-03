package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestToBase64 File/Global.go File/ToBase64.go File/ToBase64_test.go File/Ext.go File/IsExist.go
func TestToBase64(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main222.go"
	res1,res2,err := toolFile.ToBase64(form)
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(err)
}

// go test -v -run TestToBase64 -bench=BenchmarkToBase64 -count=5 File/Global.go File/ToBase64.go File/ToBase64_test.go
func BenchmarkToBase64(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	for i:=0; i< t.N; i++ {
		_,_,_ = toolFile.ToBase64(form)
	}
}