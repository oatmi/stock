package lib

import (
	"encoding/json"
	"time"
)

// CurrentDate 返回当前时间的日期
//
// e.g. 2006-01-02 15:04:05
func CurrentDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func TimestampToDate(ts int64) string {
	return time.Unix(ts, 0).Format("2006-01-02 15:04:05")
}

// ByteToType 把字节转换成目标结构
func ConvertTo(in, out interface{}) error {
	bt, err := json.Marshal(in)
	if err != nil {
		return err
	}
	return json.Unmarshal(bt, &out)
}
