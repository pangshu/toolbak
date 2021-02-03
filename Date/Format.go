package Date

import (
	"strings"
	"time"
)

// Date 格式化时间.
// format 格式,如"Y-m-d H:i:s".
// ts为int/int64类型时间戳或time.Time类型.
func (*Date)Format(format string, ts ...interface{}) string {
	replacer := strings.NewReplacer(DatePatterns...)
	format = replacer.Replace(format)

	var t time.Time
	if len(ts) > 0 {
		val := ts[0]
		if v, ok := val.(time.Time); ok {
			t = v
		} else if v, ok := val.(int); ok {
			t = time.Unix(int64(v), 0)
		} else if v, ok := val.(int64); ok {
			t = time.Unix(int64(v), 0)
		} else {
			return ""
		}
	} else {
		t = time.Now()
	}

	return t.Format(format)
}
