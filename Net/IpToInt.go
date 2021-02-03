package Net

import (
	"encoding/binary"
	"net"
)

// IpToInt 将 IPv4 字符串形式转为 uint32
func (*Net)IpToInt(ipStr string) uint32 {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}
