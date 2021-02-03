package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestJsonToMap Convert/Global.go Convert/JsonToMap.go Convert/JsonToMap_test.go
func TestJsonToMap(t *testing.T) {
	var toolConvert Convert
	str := `{"Name":"Bob","Age":32,"Company":"IBM","Engineer":true}`
	res1,_ := toolConvert.JsonToMap(str)
	fmt.Println(res1)
}

// go test -v -run TestJsonToMap -bench=BenchmarkJsonToMap -count=5 Convert/Global.go Convert/JsonToMap.go Convert/JsonToMap_test.go
func BenchmarkJsonToMap(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := `[{"Name":"Bob","Age":32,"Company":"IBM","Engineer":true},{"Name":"John","Age":20,"Company":"Oracle","Engineer":false},{"Name":"Henry","Age":45,"Company":"Microsoft","Engineer":false}]`
	for i:=0; i< t.N; i++ {
		_,_ = toolConvert.JsonToMap(str)
	}
}