package Date

import (
	"regexp"
	"strings"
)

// IsDate2time 检查字符串是否日期格式,并转换为时间戳.注意,时间戳可能为负数(小于1970年时).
// 匹配如:
//	0000
//	0000-00
//	0000/00
//	0000-00-00
//	0000/00/00
//	0000-00-00 00
//	0000/00/00 00
//	0000-00-00 00:00
//	0000/00/00 00:00
//	0000-00-00 00:00:00
//	0000/00/00 00:00:00
// 等日期格式.
func (toolDate *Date)IsDateToTime(str string) (bool, int64) {
	if str == "" {
		return false, 0
	} else if strings.ContainsRune(str, '/') {
		str = strings.Replace(str, "/", "-", -1)
	}

	chk := regexp.MustCompile(`^[0-9]{4}(|\-[0-9]{2}(|\-[0-9]{2}(|\s+[0-9]{2}(|:[0-9]{2}(|:[0-9]{2})))))$`).MatchString(str)
	if !chk {
		return false, 0
	}

	leng := len(str)
	if leng < 19 {
		reference := "1970-01-01 00:00:00"
		str = str + reference[leng:19]
	}

	tim, err := toolDate.StrToTimestamp(str)
	if err != nil {
		return false, 0
	}

	return true, tim
}
