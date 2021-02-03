package Url

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestRawEncode Url/Global.go Url/RawEncode.go Url/RawEncode_test.go
func TestRawEncode(t *testing.T) {
	var toolUrl Url
	str := `foo bar@baz`
	res1 := toolUrl.RawEncode(str)
	fmt.Println(res1)
}

// go test -v -run TestRawEncode -bench=BenchmarkRawEncode -count=5 Url/Global.go Url/RawEncode.go Url/RawEncode_test.go  Url/IsUrl.go
func BenchmarkRawEncode(t *testing.B) {
	t.ResetTimer()
	str := `foo bar@baz`
	var toolUrl Url
	for i:=0; i< t.N; i++ {
		_ = toolUrl.RawEncode(str)
	}
}