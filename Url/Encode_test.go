package Url

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestEncode Url/Global.go Url/Encode.go Url/Encode_test.go
func TestEncode(t *testing.T) {
	var toolUrl Url
	str := "one & two"
	res1 := toolUrl.Encode(str)
	fmt.Println(res1)
}

// go test -v -run TestEncode -bench=BenchmarkEncode -count=5 Url/Global.go Url/Encode.go Url/Encode_test.go
func BenchmarkEncode(t *testing.B) {
	t.ResetTimer()
	str := "one & two"
	var toolUrl Url
	for i:=0; i< t.N; i++ {
		_ = toolUrl.Encode(str)
	}
}