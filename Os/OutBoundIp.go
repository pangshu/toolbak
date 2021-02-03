package Os

import "net"

// 出口IP
func (*Os)OutBoundIp() string {
	res := ""
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	if conn != nil {
		addr := conn.LocalAddr().(*net.UDPAddr)
		res = addr.IP.String()
		_ = conn.Close()
	}

	return res
}
