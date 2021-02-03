package Net

//import (
//	"net"
//	"strings"
//)
//
//// IsPortOpen 检查主机端口是否开放.protocols为协议名称,可选,默认tcp.
//func IsPortOpen(host string, port interface{}, protocols ...string) bool {
//	// 默认tcp协议
//	protocol := "tcp"
//	if len(protocols) > 0 && len(protocols[0]) > 0 {
//		protocol = strings.ToLower(protocols[0])
//	}
//
//	conn, _ := net.DialTimeout(protocol, net.JoinHostPort(host, KConv.ToStr(port)), CHECK_CONNECT_TIMEOUT)
//	if conn != nil {
//		_ = conn.Close()
//		return true
//	}
//
//	return false
//}
