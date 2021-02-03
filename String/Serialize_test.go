package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestSerialize String/Global.go String/Serialize.go String/Serialize_test.go
func TestSerialize(t *testing.T) {
	var toolbaktring String
	res1,_ := toolbaktring.Serialize("aa bb cc AaBbCcDd")
	fmt.Println(res1)
}

// go test -v -run TestSerialize -bench=BenchmarkSerialize -count=5 String/Global.go String/Serialize.go String/Serialize_test.go
func BenchmarkSerialize(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_,_ = toolbaktring.Serialize("aa bb cc AaBbCcDd")
	}
}