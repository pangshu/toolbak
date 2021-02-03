package Url

import (
	"fmt"
	"strings"
	"testing"
)

// 测试命令: go test -v -run TestFormatUrl Url/FormatUrl_test.go Url/Global.go Url/FormatUrl.go
func TestFormatUrl(t *testing.T) {
	var toolUrl Url
	res1 := toolUrl.FormatUrl("")
	res2 := toolUrl.FormatUrl("abc.com",2)
	res3 := toolUrl.FormatUrl("abc.com/hello?a=1",1)
	res4 := toolUrl.FormatUrl(`www.aaa.abc.com:8080\/a//b/c///d\\12/33\.jpg`)
	res5 := toolUrl.FormatUrl("/abc.com/hello?a=1",1)
	res6 := toolUrl.FormatUrl("https://abc.com/hello?a=1",1)
	res7 := toolUrl.FormatUrl(`/a//b/c///d\\12/33.jpg`)
	if res1 != "" || res2 == "" || res3 == "" || strings.ContainsRune(res4, '\\') || res5 == "" || res6 == "" || res7 == "" {
		t.Errorf("FormatUrl fail")
	}
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println(res4)
	fmt.Println(res5)
	fmt.Println(res6)
	fmt.Println(res7)
}

// go test -v -run TestFormatUrl -bench=BenchmarkFormatUrl -count=5 Url/FormatUrl_test.go Url/Global.go Url/FormatUrl.go
func BenchmarkFormatUrl(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	for i:=0; i< t.N; i++ {
		var toolUrl Url
		res1 := toolUrl.FormatUrl("")
		res2 := toolUrl.FormatUrl("abc.com",2)
		res3 := toolUrl.FormatUrl("abc.com/hello?a=1",1)
		res4 := toolUrl.FormatUrl(`www.aaa.abc.com:8080\/a//b/c///d\\12/33.jpg`)
		res5 := toolUrl.FormatUrl("/abc.com/hello?a=1",1)
		res6 := toolUrl.FormatUrl("https://abc.com/hello?a=1",1)
		res7 := toolUrl.FormatUrl(`/a//b/c///d\\12/33.jpg`)
		if res1 != "" || res2 == "" || res3 == "" || strings.ContainsRune(res4, '\\') || res5 == "" || res6 == "" || res7 == "" {
			t.Errorf("FormatUrl fail")
		}
	}
}