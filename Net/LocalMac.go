package Net

import "net"

// 获取本机MAC地址
func (*Net)GetLocalMac() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for _, inter := range interfaces {
		address, err := inter.Addrs()
		if err != nil {
			return ""
		}

		for _, address := range address {
			// check the address type and if it is not a loopback the display it
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					return inter.HardwareAddr.String()
				}
			}
		}
	}
	return ""
}
