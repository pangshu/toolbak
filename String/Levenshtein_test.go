package String

import (
	"fmt"
	"strings"
	"testing"
)

// 测试命令: go test -v -run TestLevenshtein String/Global.go String/Levenshtein.go String/Levenshtein_test.go
func TestLevenshtein(t *testing.T) {
	var toolbaktring String
	s1 := "frederick"
	s2 := "fredelstick"

	res1 := toolbaktring.Levenshtein(&s1, &s2)
	res2 := toolbaktring.Levenshtein(&s2, &s1)
	res3 := toolbaktring.Levenshtein(&s1, &s1)
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	if res1 != res2 || res3 != 0 {
		t.Error("Levenshtein fail")
		return
	}

	s3 := "中国"
	s4 := "中华人民共和国"
	s5 := "中华"
	s6 := ""
	s7 := strings.Repeat(s4, 15)
	res4 := toolbaktring.Levenshtein(&s3, &s4)
	res5 := toolbaktring.Levenshtein(&s4, &s5)
	res6 := toolbaktring.Levenshtein(&s5, &s6)
	res7 := toolbaktring.Levenshtein(&s5, &s7)
	fmt.Println(res4)
	fmt.Println(res5)
	fmt.Println(res6)
	fmt.Println(res7)

	if res4 != res5 || res6 <= 0 || res7 != -1 {
		t.Error("Levenshtein fail")
		return
	}
}

// go test -v -run TestLevenshtein -bench=BenchmarkLevenshtein -count=5 String/Global.go String/Levenshtein.go String/Levenshtein_test.go
func BenchmarkLevenshtein(t *testing.B) {
	t.ResetTimer()
	var toolbaktring String
	s1 := "frederick"
	s2 := "fredelstick"
	for i:=0; i< t.N; i++ {
		_ = toolbaktring.Levenshtein(&s1, &s2)
	}
}