package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestExt File/Global.go File/Ext.go File/Ext_test.go File/IsExist.go
func TestExt(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/core/config/aaa.sql"
	res1 := toolFile.Ext(form)
	fmt.Println(res1)
}

// go test -v -run TestExt -bench=BenchmarkExt -count=5 File/Global.go File/Ext.go File/Ext_test.go
func BenchmarkExt(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v5/core/router"
	for i:=0; i< t.N; i++ {
		_ = toolFile.Ext(form)
	}
}