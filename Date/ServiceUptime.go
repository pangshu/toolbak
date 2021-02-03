package Date

import "time"

// ServiceUptime 获取当前服务运行时间,纳秒int64.
func (*Date)ServiceUptime() time.Duration {
	return time.Since(time.Now())
}
