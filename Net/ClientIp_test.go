package Net

import (
	"fmt"
	"net/http"
	"testing"
)

// 测试命令: go test -v -run TestGetClientIp Net/Global.go Net/GetClientIp.go Net/GetClientIp_test.go
func TestGetClientIp(t *testing.T) {
	var toolNet Net
	newRequest := func(remoteAddr, xRealIP string, xForwardedFor ...string) *http.Request {
		h := http.Header{}
		h.Set("X-Real-IP", xRealIP)
		for _, address := range xForwardedFor {
			h.Set("X-Forwarded-For", address)
		}

		return &http.Request{
			RemoteAddr: remoteAddr,
			Header:     h,
		}
	}

	c := newRequest("114.114.114.114", "8.8.8.8")
	res1 := toolNet.GetClientIp(c)
	fmt.Println(res1)
}

// go test -v -run TestGetClientIp -bench=BenchmarkGetClientIp -count=5 Net/Global.go Net/GetClientIp.go Net/GetClientIp_test.go
func BenchmarkGetClientIp(t *testing.B) {
	t.ResetTimer()
	var toolNet Net
	newRequest := func(remoteAddr, xRealIP string, xForwardedFor ...string) *http.Request {
		h := http.Header{}
		h.Set("X-Real-IP", xRealIP)
		for _, address := range xForwardedFor {
			h.Set("X-Forwarded-For", address)
		}

		return &http.Request{
			RemoteAddr: remoteAddr,
			Header:     h,
		}
	}

	for i:=0; i< t.N; i++ {
		c := newRequest("114.114.114.114", "8.8.8.8")
		_ = toolNet.GetClientIp(c)
	}
}