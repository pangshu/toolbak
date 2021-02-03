package Os

import (
	"net"
)

// GetIpByHostname 返回主机名对应的 IPv4地址.
func (*Os)GetIpByHostname(hostname string) (string, error) {
	ips, err := net.LookupIP(hostname)
	if ips != nil {
		for _, v := range ips {
			if v.To4() != nil {
				return v.String(), nil
			}
		}
		return "", nil
	}
	return "", err
}
