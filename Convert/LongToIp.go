package Convert

import (
	"encoding/binary"
	"net"
)

// LongToIp 将长整型转化为字符串形式带点的互联网标准格式地址(IPV4).
func (*Convert)LongToIp(properAddress uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, properAddress)
	ip := net.IP(ipByte)
	return ip.String()
}
