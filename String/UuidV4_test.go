package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestUuidV4 String/Global.go String/UuidV4.go String/UuidV4_test.go
func TestUuidV4(t *testing.T) {
	var toolbaktring String
	res1,_ := toolbaktring.UuidV4()
	fmt.Println(res1)
}

// go test -v -run TestUuidV4 -bench=BenchmarkUuidV4 -count=5 String/Global.go String/UuidV4.go String/UuidV4_test.go
func BenchmarkUuidV4(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_,_ = toolbaktring.UuidV4()
	}
}