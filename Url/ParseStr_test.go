package Url

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestParseStr Url/Global.go Url/ParseStr.go Url/ParseStr_test.go
func TestParseStr(t *testing.T) {
	var toolUrl Url
	str := `f[a][a]=m&f[a][b]=n`	//f[a][]=1&f[a][]=c&f[a][]=&f[b][]=bb&f[]=3&f[]=4
	arr1 := make(map[string]interface{})
	res1 := toolUrl.ParseStr(str, arr1)
	fmt.Println(res1)
}

// go test -v -run TestParseStr -bench=BenchmarkParseStr -count=5 Url/Global.go Url/ParseStr.go Url/ParseStr_test.go  Url/IsUrl.go
func BenchmarkParseStr(t *testing.B) {
	t.ResetTimer()
	str := `first=value&arr[]=foo+bar&arr[]=baz`
	arr1 := make(map[string]interface{})
	var toolUrl Url
	for i:=0; i< t.N; i++ {
		_ = toolUrl.ParseStr(str, arr1)
	}
}