package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsExist File/Global.go File/IsExist.go File/IsExist_test.go
func TestIsExist(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	res1 := toolFile.IsExist(form)
	fmt.Println(res1)
}

// go test -v -run TestIsExist -bench=BenchmarkIsExist -count=5 File/Global.go File/IsExist.go File/IsExist_test.go
func BenchmarkIsExist(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v4"
	for i:=0; i< t.N; i++ {
		_ = toolFile.IsExist(form)
	}
}