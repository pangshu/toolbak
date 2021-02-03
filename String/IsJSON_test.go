package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsJSON String/Global.go String/IsJSON.go String/IsJSON_test.go
func TestIsJSON(t *testing.T) {
	var toolbaktring String
	res1 := toolbaktring.IsJSON(`{"meta":{"Status":200,"Content-Type":"application/json","Content-Length":"19","etc":"etc"},"data":{"name":"yummy"}}`)
	fmt.Println(res1)
}

// go test -v -run TestIsJSON -bench=BenchmarkIsJSON -count=5 String/Global.go String/IsJSON.go String/IsJSON_test.go
func BenchmarkIsJSON(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolbaktring String
		_ = toolbaktring.IsJSON(`{"meta":{"Status":200,"Content-Type":"application/json","Content-Length":"19","etc":"etc"},"data":{"name":"yummy"}}`)
	}
}