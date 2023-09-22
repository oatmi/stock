package lib

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
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

// NewBatchNO 入库批次号生成
//
// 当前日期数字表示
func NewBatchNO() string {
	return time.Now().Format("20060102150405")
}

// UserName 从Cookie获取用户名
func UserName(c *gin.Context) string {
	name, err := c.Cookie("stock_un")
	if err != nil {
		return ""
	}
	return name
}
