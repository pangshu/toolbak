package Date

import "time"

// Time 获取当前Unix时间戳(秒).
func (*Date)UnixTime() int64 {
	return time.Now().Unix()
}
