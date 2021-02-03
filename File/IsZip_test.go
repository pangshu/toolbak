package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsZip File/Global.go File/IsZip.go File/IsZip_test.go File/Ext.go File/IsExist.go
func TestIsZip(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	res1 := toolFile.IsZip(form)
	fmt.Println(res1)
}

// go test -v -run TestIsZip -bench=BenchmarkIsZip -count=5 File/Global.go File/IsZip.go File/IsZip_test.go
func BenchmarkIsZip(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	for i:=0; i< t.N; i++ {
		_ = toolFile.IsZip(form)
	}
}