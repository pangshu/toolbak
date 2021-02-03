package Url

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestParseUrl Url/Global.go Url/ParseUrl.go Url/ParseUrl_test.go
func TestParseUrl(t *testing.T) {
	var toolUrl Url
	str := `https://www.cnbeta.com/articles/tech/1076401.htm`
	res1,_ := toolUrl.ParseUrl(str, -1)
	fmt.Println(res1)
}

// go test -v -run TestParseUrl -bench=BenchmarkParseUrl -count=5 Url/Global.go Url/ParseUrl.go Url/ParseUrl_test.go  Url/IsUrl.go
func BenchmarkParseUrl(t *testing.B) {
	t.ResetTimer()
	str := `https://www.cnbeta.com/articles/tech/1076401.htm`
	var toolUrl Url
	for i:=0; i< t.N; i++ {
		_,_ = toolUrl.ParseUrl(str, -1)
	}
}