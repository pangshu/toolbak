package Url

import (
	"fmt"
	"net/url"
	"testing"
)

// 测试命令: go test -v -run TestBuildQuery Url/Global.go Url/BuildQuery.go Url/BuildQuery_test.go
func TestBuildQuery(t *testing.T) {
	var toolUrl Url
	params := url.Values{}
	params.Add("a", "abc")
	params.Add("b", "123")
	params.Add("c", "你好")
	res1 := toolUrl.BuildQuery(params)
	fmt.Println(res1)
}

// go test -v -run TestBuildQuery -bench=BenchmarkBuildQuery -count=5 Url/Global.go Url/BuildQuery.go Url/BuildQuery_test.go
func BenchmarkBuildQuery(t *testing.B) {
	t.ResetTimer()
	params := url.Values{}
	params.Add("a", "abc")
	params.Add("b", "123")
	params.Add("c", "你好")
	var toolUrl Url
	for i:=0; i< t.N; i++ {
		_ = toolUrl.BuildQuery(params)
	}
}