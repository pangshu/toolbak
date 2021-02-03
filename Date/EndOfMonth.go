package Date

import "time"

// EndOfMonth 获取日期中当月的结束时间.
func (toolDate *Date)EndOfMonth(date time.Time) time.Time {
	return toolDate.StartOfMonth(date).AddDate(0, 1, 0).Add(-time.Nanosecond)
}
