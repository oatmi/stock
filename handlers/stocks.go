package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oatmi/stock/data"
	"github.com/oatmi/stock/data/sqlite"
	"github.com/spf13/cast"
)

type AisudaiResponse struct {
	Status  int         `json:"status"` // 状态码，为0表示成功
	Message string      `json:"msg"`    // 文案信息
	Data    interface{} `json:"data"`   // 返回值内容
}

type AisudaiCRUDData struct {
	Count int         `json:"count"`
	Rows  interface{} `json:"raws"`
}

// GetStocks 获取库存数据
func GetStocks(c *gin.Context) {
	query := sqlite.New(data.Sqlite3)

	list, err := query.ListStocks(c, buildListStockParams(c))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	count, err := query.CountStocks(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, AisudaiCRUDData{Count: int(count), Rows: list})
	return
}

func buildListStockParams(c *gin.Context) sqlite.ListStocksParams {
	arg := sqlite.ListStocksParams{}
	if val, ok := c.GetQuery("Name"); ok {
		arg.Name = "%" + val + "%"
	}
	if val, ok := c.GetQuery("ProductType"); ok {
		arg.ProductType = cast.ToInt64(val)
	}
	if val, ok := c.GetQuery("Type"); ok {
		arg.Type = cast.ToInt64(val)
	}
	if val, ok := c.GetQuery("ProduceDate"); ok {
		arg.ProduceDate = val
	}
	if val, ok := c.GetQuery("Location"); ok {
		arg.Location = val
	}
	fmt.Printf("debug: %+v\n", arg)
	return arg
}
