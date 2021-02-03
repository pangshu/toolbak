package Date

import "time"

// DaysBetween 获取两个日期的间隔天数.
func (*Date)DaysBetween(fromDate, toDate time.Time) int {
	return int(toDate.Sub(fromDate) / (24 * time.Hour))
}
