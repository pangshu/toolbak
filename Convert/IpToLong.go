package Convert

import (
	"encoding/binary"
	"net"
)

// IpToLong 将 IPV4 的字符串互联网协议转换成长整型数字.
func (*Convert)IpToLong(ipAddress string) uint32 {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return 0
	}
	return binary.BigEndian.Uint32(ip.To4())
}
