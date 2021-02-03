package String

import (
	"testing"
)

// 测试命令: go test -v -run TrimBom String/Global.go String/TrimBom.go String/TrimBom_test.go
func TestTrimBom(t *testing.T) {
	var toolbaktring String
	tests := []struct {
		str      []byte
		expected string
	}{
		{[]byte{}, ""},
		{[]byte("hello"), "hello"},
		{[]byte("\xEF\xBB\xBF"), ""},
		{[]byte("\xef\xbb\xbf"), ""},
		{[]byte("\xEF\xBB\xBFhello"), "hello"},
		{[]byte("\xEF\xBB\xBFworld"), "world"},
	}
	for _, test := range tests {
		actual := toolbaktring.TrimBom(test.str)
		if string(actual) != test.expected {
			t.Errorf("Expected TrimBOM(%v) to be %v, got %v", test.str, test.expected, actual)
			return
		}
	}
}

// go test -v -run TrimBom -bench=BenchmarkTrimBom -count=5 String/Global.go String/TrimBom.go String/TrimBom_test.go
func BenchmarkTrimBom(t *testing.B) {
	t.StopTimer()
	t.StartTimer()
	var toolbaktring String
	str := []byte("\xEF\xBB\xBFhello")
	for i := 0; i < t.N; i++ {
		toolbaktring.TrimBom(str)
	}
}