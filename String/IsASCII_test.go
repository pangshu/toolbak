package String

import (
	"fmt"
	"testing"
)

// 测试命令: go test -v -run TestIsASCII String/Global.go String/IsASCII.go String/IsASCII_test.go
func TestIsASCII(t *testing.T) {
	var toolbaktring String
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"ｆｏｏbar", false},
		{"ｘｙｚ０９８", false},
		{"１２３456", false},
		{"你好，世界", false},
		{"foobar", true},
		{"0987654321", true},
		{"test@example.com", true},
		{"1234abcDEF", true},
	}
	for _, test := range tests {
		actual := toolbaktring.IsASCII(test.param)
		fmt.Println(actual)
		if actual != test.expected {
			t.Errorf("Expected IsASCII(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

// go test -v -run TestIsASCII -bench=BenchmarkIsASCII -count=5 String/Global.go String/IsASCII.go String/IsASCII_test.go
func BenchmarkIsASCII(t *testing.B) {
	t.ResetTimer()
	var toolbaktring String
	for i:=0; i< t.N; i++ {
		_ = toolbaktring.IsASCII("Who's Bill Gates?")
	}
}