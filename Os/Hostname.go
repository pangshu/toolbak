package Os

import "os"

// Hostname 获取主机名.
func (*Os)Hostname() (string, error) {
	return os.Hostname()
}
