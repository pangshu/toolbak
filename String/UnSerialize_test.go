package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestUnSerialize String/Global.go String/Serialize.go String/UnSerialize.go String/UnSerialize_test.go
func TestUnSerialize(t *testing.T) {
	var toolbaktring String
	res1,_ := toolbaktring.Serialize("aa bb cc AaBbCcDd")
	res2,_ := toolbaktring.UnSerialize(res1)
	fmt.Println(res1)
	fmt.Println(res2)
}

// go test -v -run TestUnSerialize -bench=BenchmarkUnSerialize -count=5 String/Global.go String/Serialize.go String/UnSerialize.go String/UnSerialize_test.go
func BenchmarkUnSerialize(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		res1,_ := toolbaktring.Serialize("aa bb cc AaBbCcDd")
		_,_ = toolbaktring.UnSerialize(res1)
	}
}