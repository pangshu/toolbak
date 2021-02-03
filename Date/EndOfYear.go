package Date

import "time"

// EndOfYear 获取日期中当年的结束时间.
func (toolDate *Date)EndOfYear(date time.Time) time.Time {
	return toolDate.StartOfYear(date).AddDate(1, 0, 0).Add(-time.Nanosecond)
}