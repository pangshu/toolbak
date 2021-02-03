package Os

import "os"

// SetEnv 设置一个环境变量的值.
func (*Os)SetEnv(name, data string) error {
	return os.Setenv(name, data)
}
