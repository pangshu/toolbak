package Os

import (
	"net"
	"strings"
)

// GetHostByIp 获取指定的IP地址对应的主机名.
func (*Os)GetHostByIp(ipAddress string) (string, error) {
	names, err := net.LookupAddr(ipAddress)
	if names != nil {
		return strings.TrimRight(names[0], "."), nil
	}
	return "", err
}
