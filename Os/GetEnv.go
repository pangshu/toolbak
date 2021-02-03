package Os

import "os"

// GetEnv 获取一个环境变量的值.def为默认值.
func (*Os)GetEnv(name string, def ...string) string {
	val := os.Getenv(name)
	if val == "" && len(def) > 0 {
		val = def[0]
	}

	return val
}
