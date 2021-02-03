package Os

import "net"

// GetIpsByHost 获取互联网域名/主机名对应的 IPv4 地址列表.
func (*Os)GetIpsByDomain(domain string) ([]string, error) {
	ips, err := net.LookupIP(domain)
	if ips != nil {
		var ipstrs []string
		for _, v := range ips {
			if v.To4() != nil {
				ipstrs = append(ipstrs, v.String())
			}
		}
		return ipstrs, nil
	}
	return nil, err
}
