package lib

import "time"

// CurrentDate 返回当前时间的日期
//
// e.g. 2006-01-02 15:04:05
func CurrentDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func TimestampToDate(ts int64) string {
	return time.Unix(ts, 0).Format("2006-01-02 15:04:05")
}
