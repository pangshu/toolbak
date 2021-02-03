package Convert

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestJsonpToJson Convert/Global.go Convert/JsonpToJson.go Convert/JsonpToJson_test.go Convert/Gettype.go
func TestJsonpToJson(t *testing.T) {
	var toolConvert Convert
	str := `myFunc([{"Name":"Bob","Age":32,"Company":"IBM","Engineer":true},{"Name":"John","Age":20,"Company":"Oracle","Engineer":false},{"Name":"Henry","Age":45,"Company":"Microsoft","Engineer":false}]);`
	res1,_ := toolConvert.JsonpToJson(str)
	fmt.Println(res1)
}

// go test -v -run TestJsonpToJson -bench=BenchmarkJsonpToJson -count=5 Convert/Global.go Convert/JsonpToJson.go Convert/JsonpToJson_test.go Convert/Gettype.go
func BenchmarkJsonpToJson(t *testing.B) {
	t.ResetTimer()
	var toolConvert Convert
	str := `myFunc([{"Name":"Bob","Age":32,"Company":"IBM","Engineer":true},{"Name":"John","Age":20,"Company":"Oracle","Engineer":false},{"Name":"Henry","Age":45,"Company":"Microsoft","Engineer":false}]);`
	for i:=0; i< t.N; i++ {
		_,_ = toolConvert.JsonpToJson(str)
	}
}