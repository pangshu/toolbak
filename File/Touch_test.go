package File

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestTouch File/Global.go File/Touch.go File/Touch_test.go File/IsDir.go
func TestTouch(t *testing.T) {
	var toolFile File
	form := "/www/golang/src/v5/main333.go"
	res1 := toolFile.Touch(form, 1)
	fmt.Println(res1)
}

// go test -v -run TestTouch -bench=BenchmarkTouch -count=5 File/Global.go File/Touch.go File/Touch_test.go
func BenchmarkTouch(t *testing.B) {
	t.ResetTimer()
	var toolFile File
	form := "/www/golang/src/v5/main.go"
	for i:=0; i< t.N; i++ {
		_ = toolFile.Touch(form, 1)
	}
}