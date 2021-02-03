package Url

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestDecode Url/Global.go Url/Decode.go Url/Decode_test.go
func TestDecode(t *testing.T) {
	var toolUrl Url
	str := "one%20%26%20two"
	res1,_ := toolUrl.Decode(str)
	fmt.Println(res1)
}

// go test -v -run TestDecode -bench=BenchmarkDecode -count=5 Url/Global.go Url/Decode.go Url/Decode_test.go
func BenchmarkDecode(t *testing.B) {
	t.ResetTimer()
	str := "one%20%26%20two"
	var toolUrl Url
	for i:=0; i< t.N; i++ {
		_,_ = toolUrl.Decode(str)
	}
}