package Date

import "time"

// Usleep 以指定的微秒数延迟执行.
func (*Date)Usleep(t int64) {
	time.Sleep(time.Duration(t) * time.Microsecond)
}
