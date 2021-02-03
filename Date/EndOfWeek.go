package Date

import "time"

// EndOfWeek 获取日期中当周的结束时间;
// weekStartDay 周几作为周的第一天,本库默认周一.
func (toolDate *Date)EndOfWeek(date time.Time, weekStartDay ...time.Weekday) time.Time {
	return toolDate.StartOfWeek(date, weekStartDay...).AddDate(0, 0, 7).Add(-time.Nanosecond)
}
