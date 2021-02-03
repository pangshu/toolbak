package Date

// Str2Timestamp 将字符串转换为时间戳,秒.
// str 为要转换的字符串;
// format 为该字符串的格式,默认为"2006-01-02 15:04:05" .
func (toolDate *Date)StrToTimestamp(str string, format ...string) (int64, error) {
	tim, err := toolDate.StrToTimestruct(str, format...)
	if err != nil {
		return 0, err
	}

	return tim.Unix(), nil
}
