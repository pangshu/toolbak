package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsImg File/Global.go File/IsImg.go File/IsImg_test.go File/IsExist.go File/Ext.go
func TestIsImg(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	res1 := toolFile.IsImg(form)
	fmt.Println(res1)
}

// go test -v -run TestIsImg -bench=BenchmarkIsImg -count=5 File/Global.go File/IsImg.go File/IsImg_test.go
func BenchmarkIsImg(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	for i:=0; i< t.N; i++ {
		_ = toolFile.IsImg(form)
	}
}